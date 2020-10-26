export class SessionUser {
    userId: string;
    name: string;
    token: string;
    isActive: boolean;
    email: string;
    language: string;
    isAdmin: boolean;
}

export class Profile {
    user: SessionUser;
    token: string;
    roleMenus: RoleMenu[] = [];
    permissions: Permission[] = [];
}

export class RoleMenu {
    projectId: string;
    projectName: string;
    menus: string[];
}

export class Permission {
    projectId: string;
    projectName: string;
    projectRole: string;
    userPermissionRoles: UserPermissionRole[] = [];
}

export class UserPermissionRole {
    operation: string;
    roles: string[];
}
