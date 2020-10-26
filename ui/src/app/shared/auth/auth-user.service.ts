import {Injectable} from '@angular/core';
import {CommonRoutes} from '../../constant/route';
import {
    ActivatedRouteSnapshot,
    CanActivate,
    CanActivateChild,
    Router,
    RouterStateSnapshot,
    UrlTree
} from '@angular/router';
import {Observable} from 'rxjs';
import {SessionService} from './session.service';
import {Profile} from './session-user';
import {AlertLevels} from '../../layout/common-alert/alert';
import {ModalAlertService} from '../common-component/modal-alert/modal-alert.service';

@Injectable({
    providedIn: 'root'
})
export class AuthUserService implements CanActivate, CanActivateChild {

    constructor(private sessionService: SessionService, private router: Router,private modalAlertService: ModalAlertService) {
    }

    canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot):
        Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
        return new Observable<boolean>((observer) => {
            this.isLogin().subscribe(res => {
                this.sessionService.cacheProfile(res);
                observer.next(true);
                observer.complete();
            }, error => {
                observer.next(false);
                observer.complete();
                this.modalAlertService.showAlert(error.error.msg, AlertLevels.ERROR);
                this.router.navigateByUrl(CommonRoutes.LOGIN).then();
            });
        });
    }

    canActivateChild(childRoute: ActivatedRouteSnapshot, state: RouterStateSnapshot):
        Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
        return this.canActivate(childRoute, state);
    }

    isLogin(): Observable<Profile> {
        return this.sessionService.getProfile();
    }
}
