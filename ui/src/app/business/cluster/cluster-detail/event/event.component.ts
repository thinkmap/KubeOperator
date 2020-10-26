import {Component, OnInit} from '@angular/core';
import {KubernetesService} from '../../kubernetes.service';
import {ActivatedRoute} from '@angular/router';
import {Cluster} from '../../cluster';
import {EventService} from './event.service';
import {CommonAlertService} from '../../../../layout/common-alert/common-alert.service';
import {TranslateService} from '@ngx-translate/core';
import {AlertLevels} from '../../../../layout/common-alert/alert';

@Component({
    selector: 'app-event',
    templateUrl: './event.component.html',
    styleUrls: ['./event.component.css']
})
export class EventComponent implements OnInit {

    loading = false;
    currentCluster: Cluster;
    namespaces;
    events;
    currentNamespace: string;
    npdExists = false;

    constructor(private kubernetesService: KubernetesService,
                private route: ActivatedRoute,
                private eventService: EventService,
                private commonAlertService: CommonAlertService,
                private translateService: TranslateService) {
    }

    ngOnInit(): void {
        this.loading = true;
        this.route.parent.data.subscribe(data => {
            this.currentCluster = data.cluster;
            this.kubernetesService.listNamespaces(this.currentCluster.name).subscribe(res => {
                this.namespaces = res.items;
                if (this.namespaces.length > 0) {
                    const namespace = this.namespaces[0];
                    this.currentNamespace = namespace.metadata.name;
                    this.listEvents(this.currentNamespace);
                }
            });
            this.getNpdExists();
        });
    }

    listEvents(namespace: string) {
        this.loading = true;
        this.kubernetesService.listEventsByNamespace(this.currentCluster.name, namespace).subscribe(res => {
            this.events = res.items;
            this.loading = false;
        });
    }

    getNpdExists() {
        this.kubernetesService.listPod(this.currentCluster.name).subscribe(data => {
            const pods = data.items;
            for (const pod of pods) {
                if (pod.metadata.generateName === 'node-problem-detector-') {
                    this.npdExists = true;
                    break;
                }
            }
        });
    }

    changeNpd(exists) {
        this.npdExists = !exists;
        let op = 'create';
        if (exists) {
            op = 'delete';
        }
        this.eventService.changeNpd(this.currentCluster.name, op).subscribe(res => {
            this.commonAlertService.showAlert(this.translateService.instant('APP_UPDATE_SUCCESS'), AlertLevels.SUCCESS);
        }, error => {
            this.commonAlertService.showAlert(error.error.msg, AlertLevels.ERROR);
        });
    }
}
