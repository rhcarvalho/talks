from django.db import models


# TODO implement a model

class Deal(models.Model):
    pass


class Deal(models.Model):
    name = models.CharField(max_length=255)
    value = models.IntegerField()
    updated_at = models.DateTimeField(auto_now=True)
