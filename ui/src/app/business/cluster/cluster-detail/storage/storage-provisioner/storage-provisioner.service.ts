import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {Observable} from 'rxjs';
import {CreateStorageProvisionerRequest, StorageProvisioner} from './storage-provisioner';

@Injectable({
    providedIn: 'root'
})
export class StorageProvisionerService {

    constructor(private http: HttpClient) {
    }

    baseUrl = '/api/v1/clusters/provisioner/{cluster_name}';

    list(clusterName: string): Observable<StorageProvisioner[]> {
        return this.http.get<StorageProvisioner[]>(this.baseUrl.replace('{cluster_name}', clusterName));
    }

    create(clusterName: string, item: CreateStorageProvisionerRequest): Observable<StorageProvisioner> {
        return this.http.post<StorageProvisioner>(this.baseUrl.replace('{cluster_name}', clusterName), item);
    }

    delete(clusterName: string, name: string): Observable<any> {
        const url = this.baseUrl.replace('{cluster_name}', clusterName) + name + '/';
        return this.http.delete<any>(url);
    }

    batch(clusterName: string, items: StorageProvisioner[]): Observable<any> {
        const url = this.baseUrl.replace('{cluster_name}', 'batch/' + clusterName);
        return this.http.post<any>(url, {items, operation: 'delete'});
    }
}


