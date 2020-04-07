# -*- coding: utf-8 -*-
#
from django.db import transaction
from django.shortcuts import get_object_or_404

from ..ctx import set_current_project
from ..models import Project


class ProjectResourceAPIMixin:
    lookup_kwargs = 'project_name'
    is_project_request = False
    project = None
    project_name = ''

    @transaction.atomic
    def dispatch(self, request, *args, **kwargs):
        if kwargs.get(self.lookup_kwargs):
            self.is_project_request = True
            self.project_name = kwargs[self.lookup_kwargs]
            self.project = self.get_project()
            set_current_project(self.project)
        return super().dispatch(request, *args, **kwargs)

    def get_project(self):
        if self.project is not None:
            return self.project
        if self.is_project_request:
            self.project = get_object_or_404(Project, name=self.project_name)
        return self.project

    def get_serializer_class(self):
        if hasattr(self, 'action') and self.action in ('list', 'retrieve') \
                and hasattr(self, 'read_serializer_class'):
            return self.read_serializer_class
        return super().get_serializer_class()

    def get_queryset(self):
        return super().get_queryset().all()
