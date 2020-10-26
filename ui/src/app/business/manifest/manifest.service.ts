import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {Observable} from "rxjs";
import {Manifest} from "./manifest";

@Injectable({
    providedIn: 'root'
})
export class ManifestService {

    constructor(private http: HttpClient) {
    }

    baseUrl = '/api/v1/manifests';

    list(): Observable<Manifest[]> {
        return this.http.get<Manifest[]>(this.baseUrl);
    }
}
