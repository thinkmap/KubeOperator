import os
from abc import ABCMeta, abstractmethod

from python_terraform import Terraform, IsNotFlagged

from cloud_provider.utils import generate_terraform_file, create_terrafrom_working_dir
from kubeoperator.settings import CLOUDS_RESOURCE_DIR


def get_cloud_client(vars):
    provider = vars.get('provider', {})
    from cloud_provider.clients.vsphere import VsphereCloudClient
    from cloud_provider.clients.openstack import OpenStackCloudClient
    if provider == 'vsphere':
        return VsphereCloudClient(vars)
    if provider == 'openstack':
        return OpenStackCloudClient(vars)
    else:
        return None


class CloudClient(metaclass=ABCMeta):
    cloud_config_path = CLOUDS_RESOURCE_DIR
    working_path = None

    def __init__(self, vars):
        self.vars = vars

    @abstractmethod
    def list_region(self):
        pass

    @abstractmethod
    def init_terraform(self, cluster):
        pass

    @abstractmethod
    def create_image(self, zone):
        pass

    def destroy_terraform(self, cluster):
        if not self.working_path:
            self.working_path = create_terrafrom_working_dir(cluster_name=cluster)
        t = Terraform(working_dir=self.working_path)
        p, _, _ = t.destroy('./', synchronous=False, no_color=IsNotFlagged, refresh=True)
        for i in p.stdout:
            print(i.decode())
        _, err = p.communicate()
        print(err.decode())
        return p.returncode == 0

    def apply_terraform(self, cluster, hosts_dict):
        if not self.working_path:
            self.working_path = create_terrafrom_working_dir(cluster_name=cluster.name)
        generate_terraform_file(self.working_path, self.cloud_config_path, cluster.plan.mixed_vars, hosts_dict)
        self.init_terraform(cluster)
        t = Terraform(working_dir=self.working_path)
        p, _, _ = t.apply('./', refresh=True, skip_plan=True, no_color=IsNotFlagged, synchronous=False)
        for i in p.stdout:
            print(i.decode())
        _, err = p.communicate()
        print(err.decode())
        return p.returncode == 0
