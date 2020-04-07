# Generated by Django 2.2.10 on 2020-02-24 07:41

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('kubeops_api', '0069_auto_20200224_0323'),
    ]

    operations = [
        migrations.AlterField(
            model_name='clusterhealthhistory',
            name='date_type',
            field=models.CharField(choices=[('DAY', 'DAY'), ('HOUR', 'HOUR')], default='HOUR', max_length=255),
        ),
        migrations.AlterField(
            model_name='itemrolemapping',
            name='role',
            field=models.CharField(choices=[('VIEWER', 'VIEWER'), ('MANAGER', 'MANAGER')], default='VIEWER', max_length=128),
        ),
    ]
