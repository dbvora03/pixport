from django.db import models
from django.core.files import File
from io import BytesIO
from PIL import Image

class Product(models.Model):

 
    title = models.CharField(max_length=255)
    description = models.CharField(max_length=500)
    dimensions = models.CharField(max_length=10)
    price = models.DecimalField(max_digits=6, decimal_places=2)
    image = models.CharField(max_length=255)
    slug = models.SlugField()

    class Meta:
        ordering = ('-title',)

    def __str__(self):
        return self.title
    
    def get_absolute_url(self):
        return f'{self.title}/'


    
