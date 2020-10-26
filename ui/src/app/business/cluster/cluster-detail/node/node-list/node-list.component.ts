import {Component, EventEmitter, Input, OnDestroy, OnInit, Output} from '@angular/core';
import {KubernetesService} from '../../../kubernetes.service';
import {V1Node} from '@kubernetes/client-node';
import {Cluster} from '../../../cluster';
import {ActivatedRoute} from '@angular/router';
import {NodeService} from "../node.service";
import {Node} from "../node";
import {CommonAlertService} from "../../../../../layout/common-alert/common-alert.service";
import {AlertLevels} from "../../../../../layout/common-alert/alert";
import {ClusterLoggerService} from "../../../cluster-logger/cluster-logger.service";

@Component({
    selector: 'app-node-list',
    templateUrl: './node-list.component.html',
    styleUrls: ['./node-list.component.css']
})
export class NodeListComponent implements OnInit, OnDestroy {

    loading = true;
    selected = [];
    items: Node[] = [];
    page = 1;
    timer;
    @Input() currentCluster: Cluster;
    @Output() openDetail = new EventEmitter<V1Node>();
    @Output() createEvent = new EventEmitter();
    @Output() statusEvent = new EventEmitter<Node>();
    @Output() deleteEvent = new EventEmitter<Node[]>();

    constructor(private service: KubernetesService, private route: ActivatedRoute,
                private nodeService: NodeService, private alertService: CommonAlertService, private loggerService: ClusterLoggerService) {
    }

    ngOnInit(): void {
        this.refresh();
    }

    ngOnDestroy(): void {
        clearInterval(this.timer);
    }

    refresh() {
        this.loading = true;
        this.selected = [];
        if (this.currentCluster.source === 'external') {
        }
        this.nodeService.list(this.currentCluster.name).subscribe(d => {
            this.items = d;
            this.loading = false;
        });
    }

    getInternalIp(item: Node) {
        let result = 'N/A';
        if (item.status === 'Running') {
            for (const addr of item.info.status.addresses) {
                if (addr.type === 'InternalIP') {
                    result = addr.address;
                }
            }
        }
        return result;
    }

    getVersion(item: Node) {
        let result = 'N/A';
        if (item.status === 'Running') {
            result = item.info.status.nodeInfo.kubeletVersion;
        }
        return result;
    }

    formatRAM(memory: string): string {
        let result = 0.0;
        if (memory.endsWith('Ki')) {
            const str = memory.substring(0, memory.indexOf('Ki'));
            result = parseFloat(str);
            result = result / (1024 * 1024);
        }
        return result.toFixed(2) + 'GB';
    }

    getRAM(item: Node) {
        let result = 'N/A';
        if (item.status === 'Running') {
            result = this.formatRAM(item.info.status.capacity['memory']);
        }
        return result;
    }

    getCpuCore(item: Node) {
        let result = 'N/A';
        if (item.status === 'Running') {
            result = item.info.status.capacity['cpu'];
        }
        return result;
    }

    getNodeRoles(item: Node): string[] {
        const roles: string[] = [];
        if (item.status === 'Running') {
            for (const key in item.info.metadata.labels) {
                if (key) {
                    switch (key) {
                        case 'node-role.kubernetes.io/master':
                            roles.push('master');
                            break;
                        case 'node-role.kubernetes.io/etcd':
                            roles.push('etcd');
                            break;
                        case 'node-role.kubernetes.io/worker':
                            roles.push('worker');
                            break;
                    }
                }
            }
        }
        return roles;
    }

    getStatus(item: Node) {
        if (item.status === 'Running') {
            return this.isNodeReady(item.info);
        }
        return item.status;
    }

    isNodeReady(n: V1Node): string {
        let result = 'NotReady';
        for (const condition of n.status.conditions) {
            if (condition.type === 'Ready') {
                if (condition.status === 'True') {
                    result = 'Ready';
                }
            }
        }
        return result;
    }

    onDetail(item: Node) {
        if (item.status === 'Running') {
            this.openDetail.emit(item.info);
        } else {
            this.alertService.showAlert('node is not ready', AlertLevels.ERROR);
        }
    }

    onCreate() {
        this.createEvent.emit();
    }

    onDelete() {
        this.deleteEvent.emit(this.selected);
    }

    onShowStatus(item: Node) {
        this.statusEvent.emit(item);
    }

    onShowLogger(item: Node) {
        this.loggerService.openLogger(this.currentCluster.name, item.name);
    }

    polling() {
        this.timer = setInterval(() => {
            this.nodeService.list(this.currentCluster.name).subscribe(data => {
                this.items = data;
            });
        }, 1000);
    }

}
