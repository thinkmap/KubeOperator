import {Component, EventEmitter, OnInit, Output} from '@angular/core';
import {BaseModelDirective} from '../../../shared/class/BaseModelDirective';
import {Host} from '../host';
import {HostService} from '../host.service';
import {ModalAlertService} from '../../../shared/common-component/modal-alert/modal-alert.service';
import {CommonAlertService} from '../../../layout/common-alert/common-alert.service';
import {TranslateService} from '@ngx-translate/core';
import {AlertLevels} from '../../../layout/common-alert/alert';

@Component({
    selector: 'app-host-delete',
    templateUrl: './host-delete.component.html',
    styleUrls: ['./host-delete.component.css']
})
export class HostDeleteComponent extends BaseModelDirective<Host> implements OnInit {

    opened = false;
    items: Host[] = [];
    @Output() deleted = new EventEmitter();

    constructor(private hostService: HostService, private modalAlertService: ModalAlertService,
                private commonAlertService: CommonAlertService, private translateService: TranslateService) {
        super(hostService);
    }

    ngOnInit(): void {
    }

    open(items) {
        this.opened = true;
        this.items = items;
    }

    onCancel() {
        this.opened = false;
    }

    onSubmit() {
        this.service.batch('delete', this.items).subscribe(data => {
            this.deleted.emit();
            this.opened = false;
            this.commonAlertService.showAlert(this.translateService.instant('APP_DELETE_SUCCESS'), AlertLevels.SUCCESS);
        }, error => {
            console.log(error.error.msg);
            this.modalAlertService.showAlert(error.error.msg, AlertLevels.ERROR);
        });
    }
}
