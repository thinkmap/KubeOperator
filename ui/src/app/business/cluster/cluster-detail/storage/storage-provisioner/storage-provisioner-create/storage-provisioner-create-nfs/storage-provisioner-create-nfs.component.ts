import {Component, EventEmitter, Input, OnInit, Output, ViewChild} from '@angular/core';
import {CreateStorageProvisionerRequest} from '../../storage-provisioner';
import {StorageProvisionerService} from '../../storage-provisioner.service';
import {Cluster} from '../../../../../cluster';
import {NgForm} from '@angular/forms';

@Component({
    selector: 'app-storage-provisioner-create-nfs',
    templateUrl: './storage-provisioner-create-nfs.component.html',
    styleUrls: ['./storage-provisioner-create-nfs.component.css']
})
export class StorageProvisionerCreateNfsComponent implements OnInit {

    constructor(private storageProvisionerService: StorageProvisionerService) {
    }

    opened = false;
    item: CreateStorageProvisionerRequest = new CreateStorageProvisionerRequest();
    @Output() created = new EventEmitter();
    @Input() currentCluster: Cluster;
    @ViewChild('nfsForm') nfsForm: NgForm;

    ngOnInit(): void {
    }

    open(item: CreateStorageProvisionerRequest) {
        this.opened = true;
        this.item = item;
    }

    reset() {
        this.item = new CreateStorageProvisionerRequest();
        this.nfsForm.resetForm();
    }

    onCancel() {
        this.opened = false;
    }

    onSubmit() {
        this.storageProvisionerService.create(this.currentCluster.name, this.item).subscribe(data => {
            this.opened = false;
            this.created.emit();
        });
    }

}
