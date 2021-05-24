from django.urls import path, include

from .views import LatestProductsList

urlpatterns = [
    path('products/', LatestProductsList.as_view())
]