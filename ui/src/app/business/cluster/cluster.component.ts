import {Component, OnInit, ViewChild} from '@angular/core';
import {ClusterCreateComponent} from './cluster-create/cluster-create.component';
import {ClusterListComponent} from './cluster-list/cluster-list.component';
import {ClusterDeleteComponent} from './cluster-delete/cluster-delete.component';
import {Cluster} from './cluster';
import {ClusterConditionComponent} from './cluster-condition/cluster-condition.component';
import {ClusterImportComponent} from './cluster-import/cluster-import.component';
import {ClusterUpgradeComponent} from './cluster-upgrade/cluster-upgrade.component';

@Component({
    selector: 'app-cluster',
    templateUrl: './cluster.component.html',
    styleUrls: ['./cluster.component.css']
})
export class ClusterComponent implements OnInit {

    constructor() {
    }

    @ViewChild(ClusterCreateComponent, {static: true})
    create: ClusterCreateComponent;

    @ViewChild(ClusterDeleteComponent, {static: true})
    delete: ClusterDeleteComponent;

    @ViewChild(ClusterConditionComponent, {static: true})
    condition: ClusterConditionComponent;

    @ViewChild(ClusterListComponent, {static: true})
    list: ClusterListComponent;

    @ViewChild(ClusterImportComponent, {static: true})
    import: ClusterImportComponent;

    @ViewChild(ClusterUpgradeComponent, {static: true})
    upgrade: ClusterUpgradeComponent;

    ngOnInit(): void {
    }

    openCreate() {
        this.create.open();
    }

    openImport() {
        this.import.open();
    }

    openDelete(items: Cluster[]) {
        this.delete.open(items);
    }

    openStatusDetail(cluster: Cluster) {
        this.condition.open(cluster);
    }

    refresh() {
        this.list.reset();
        this.list.pageBy();
    }

    openUpgrade(item: Cluster) {
        this.upgrade.open(item);
    }
}
