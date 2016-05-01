from rest_framework import viewsets

from api.models import Deal
from api.serializers import DealSerializer


# TODO implement a viewset

class DealViewSet(viewsets.ModelViewSet):
    pass


class DealViewSet(viewsets.ModelViewSet):
    queryset = Deal.objects.all()
    serializer_class = DealSerializer

# --------------------------------------------------------------------

from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import authentication, permissions

class DealList(APIView):
    authentication_classes = (authentication.TokenAuthentication,)
    permission_classes = (permissions.IsAdminUser,)

    def get(self, request, format=None):
        deals = ...
        return Response(...)


class DealViewSet(viewsets.ViewSet):
    def create(self, request):
        serializer = DealSerializer(data=request.DATA)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=status.HTTP_201_CREATED)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)
