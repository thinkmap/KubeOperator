# Generated by Django 2.2.10 on 2020-02-24 03:23

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('kubeops_api', '0068_auto_20200224_0305'),
    ]

    operations = [
        migrations.AlterField(
            model_name='clusterhealthhistory',
            name='date_type',
            field=models.CharField(choices=[('HOUR', 'HOUR'), ('DAY', 'DAY')], default='HOUR', max_length=255),
        ),
        migrations.AlterField(
            model_name='itemrolemapping',
            name='role',
            field=models.CharField(choices=[('VIEWER', 'VIEWER'), ('MANAGER', 'VIEWER')], default='VIEWER', max_length=128, unique=True),
        ),
    ]
