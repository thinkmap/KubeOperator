import {BaseModel, BaseRequest} from '../../../shared/class/BaseModel';
import {Region} from '../region/region';

export class Zone extends BaseModel {
    id: string;
    name: string;
    vars: string;
    credentialId: string;
    cloudVars: {} = {};
    regionName: string;
    provider: string;
    status: string;
}

export class ZoneCreateRequest extends BaseRequest {
    vars: string;
    regionName: string;
    regionID: string;
    cloudVars: {} = {};
    provider: string;
    credentialId: string;
}

export class ZoneUpdateRequest extends BaseRequest {
    vars: string;
    regionID: string;
    cloudVars: {} = {};
}

export class CloudZoneRequest extends BaseRequest {
    cloudVars: {} = {};
    datacenter: string;
}

export class CloudZone {
    cluster: string;
    networks: [] = [];
    resourcePools: [] = [];
    datastores: [] = [];
    storages: Storage[] = [];
    securityGroups: [] = [];
    networkList: Network[] = [];
    floatingNetworkList: Network[] = [];
    ipTypes: [] = [];
    imageList: Image[] = [];
}

export class CloudTemplate {
    imageName: string;
    guestId: string;
}

export class Storage {
    id: string;
    name: string;
}

export class Network {
    id: string;
    name: string;
    subnetList: Subnet[] = [];
}

export class Subnet {
    id: string;
    name: string;
}

export class Image {
    id: string;
    name: string;
}




