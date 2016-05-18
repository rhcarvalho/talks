// From https://github.com/openshift/source-to-image/blob/master/pkg/scripts/download.go

package scripts

import (
	"io"
	"net/url"
	"os"

	"github.com/golang/glog"
	"github.com/openshift/source-to-image/pkg/api"
	"github.com/openshift/source-to-image/pkg/errors"
)

type schemeReader interface {
	Read(*url.URL) (io.ReadCloser, error)
}

type downloader struct {
	schemeReaders map[string]schemeReader
}

// START OMIT
func (d *downloader) Download(url *url.URL, targetFile string) (*api.SourceInfo, error) {
	schemeReader := d.schemeReaders[url.Scheme]
	info := &api.SourceInfo{}
	if schemeReader == nil {
		glog.Errorf("No URL handler found for %s", url.String())
		return nil, errors.NewURLHandlerError(url.String())
	}

	reader, err := schemeReader.Read(url) // HL
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	// END OMIT
	out, err := os.Create(targetFile)
	defer out.Close()

	if err != nil {
		glog.Errorf("Unable to create target file %s (%s)", targetFile, err)
		return nil, err
	}

	if _, err = io.Copy(out, reader); err != nil {
		os.Remove(targetFile)
		glog.Warningf("Skipping file %s due to error copying from source: %s", targetFile, err)
		return nil, err
	}

	glog.V(2).Infof("Downloaded '%s'", url.String())
	info.Location = url.String()
	return info, nil
}
