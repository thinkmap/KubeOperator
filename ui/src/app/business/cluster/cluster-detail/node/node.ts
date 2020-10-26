import {V1Node} from "@kubernetes/client-node";
import {BaseModel} from "../../../../shared/class/BaseModel";

export class Node extends BaseModel {
    name: string;
    status: string;
    message: string;
    info: V1Node;
}


export class NodeBatch {
    hosts: string[] = [];
    nodes: string[] = [];
    increase: number;
    operation: string;
}
