from django.shortcuts import render
from rest_framework import  status

# Create your views here.
from .serializers import ProductSerializer
from rest_framework.views import APIView
from rest_framework.response import Response

from .models import Product

class LatestProductsList(APIView):
    # localhost:8000/api/v1/products
    def get(self, request, format=None):
        products = Product.objects.all()
        serializer = ProductSerializer(products, many=True)
        return Response(serializer.data)

    # localhost:8000/api/v1/products
    def post(self, request):
        serializer = ProductSerializer(data=request.data)
        serializer.is_valid(raise_exception=True)
        serializer.save()
        return Response(serializer.data, status=status.HTTP_201_CREATED)



