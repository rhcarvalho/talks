from rest_framework import routers

from api import views


router = routers.DefaultRouter()
router.register(r'v1/deals', views.DealsViewSet)

urlpatterns = router.urls
