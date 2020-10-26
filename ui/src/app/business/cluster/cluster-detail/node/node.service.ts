import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {Node, NodeBatch} from "./node";

@Injectable({
    providedIn: 'root'
})
export class NodeService {

    constructor(private http: HttpClient) {
    }

    baseUrl = '/api/v1/clusters/node/{clusterName}';
    batchUrl = '/api/v1/clusters/node/batch/{clusterName}';

    list(clusterName: string): Observable<Node[]> {
        return this.http.get<Node[]>(this.baseUrl.replace('{clusterName}', clusterName));
    }

    batch(clusterName: string, item: NodeBatch): Observable<Node[]> {
        return this.http.post<Node[]>(this.batchUrl.replace('{clusterName}', clusterName), item);
    }
}
