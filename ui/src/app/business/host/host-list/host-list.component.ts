import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import {HostService} from '../host.service';
import {BaseModelDirective} from '../../../shared/class/BaseModelDirective';
import {Host} from '../host';

@Component({
    selector: 'app-host-list',
    templateUrl: './host-list.component.html',
    styleUrls: ['./host-list.component.css']
})
export class HostListComponent extends BaseModelDirective<Host> implements OnInit {

    @Output() detailEvent = new EventEmitter<Host>();
    @Output() statusDetailEvent = new EventEmitter<Host>();
    timer;

    constructor(private hostService: HostService) {
        super(hostService);
    }

    ngOnInit(): void {
        super.ngOnInit();
        this.polling();
    }

    onDetail(item) {
        this.detailEvent.emit(item);
    }

    onStatusDetail(item: Host) {
        this.statusDetailEvent.emit(item);
    }

    polling() {
        this.timer = setInterval(() => {
            let flag = false;
            const needPolling = ['Initializing', 'Creating'];
            for (const item of this.items) {
                if (needPolling.indexOf(item.status) !== -1) {
                    flag = true;
                    break;
                }
            }
            if (flag) {
                this.hostService.list().subscribe(data => {
                    data.items.forEach(n => {
                        this.items.forEach(item => {
                            if (item.name === n.name) {
                                if (item.status !== n.status) {
                                    item.name = n.name;
                                    item.ip = n.ip;
                                    item.port = n.port;
                                    item.os = n.os;
                                    item.osVersion = n.osVersion;
                                    item.memory = n.memory;
                                    item.cpuCore = n.cpuCore;
                                    item.gpuNum = n.gpuNum;
                                    item.gpuInfo = n.gpuInfo;
                                    item.status = n.status;
                                    item.volumes = n.volumes;
                                    item.clusterName = n.clusterName;
                                }
                            }
                        });
                    });
                });
            }
        }, 5000);
    }


}
