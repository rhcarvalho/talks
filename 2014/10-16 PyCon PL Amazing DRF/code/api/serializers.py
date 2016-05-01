from rest_framework import serializers

from api.models import Deal


# TODO implement a serializer
# http://www.django-rest-framework.org/api-guide/serializers
#
# class DealSerializer(...):
#    ...


class DealSerializer(serializers.ModelSerializer):
    class Meta:
        model = Deal


class DealSerializer(serializers.Serializer):
    id = serializers.IntegerField(read_only=True)
    name = serializers.CharField(max_length=255)
    value = serializers.IntegerField()
    updated_at = serializers.DateTimeField(read_only=True)

    def restore_object(self, attrs, instance=None):
        if instance is not None:
            instance.name = attrs.get('name', instance.name)
            instance.value = attrs.get('value', instance.value)
            return instance
        return Deal(**attrs)
