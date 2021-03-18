CREATE TABLE WORKSPACES
(
    id          varchar(36) PRIMARY KEY,
    name        varchar(256) NOT NULL,
    matcher_url varchar(256),
    deleted_at  timestamp
);

CREATE TABLE AUDITIONS
(
    id            varchar(400) primary key,
    username      varchar(200)                        not null,
    table_name    varchar(200)                        not null,
    operation     varchar(200)                        not null,
    entity_id     varchar(400)                        not null,
    current_state json,
    user_ip_addr  varchar(200),
    user_agent    varchar(200),
    created_at    timestamp DEFAULT clock_timestamp() NOT NULL
);

CREATE TABLE CIRCLES
(
    id           varchar(36) PRIMARY KEY,
    name         varchar(256) NOT NULL,
    rules        jsonb        NOT NULL,
    workspace_id varchar(36)  NOT NULL,
    deleted_at   timestamp,
    CONSTRAINT fk_workspace_id_circle FOREIGN KEY (workspace_id) REFERENCES WORKSPACES (ID)
);

CREATE TABLE COMPONENTS
(
    id         varchar(36) PRIMARY KEY,
    name       varchar(256) NOT NULL,
    helm       varchar(256) NOT NULL,
    deleted_at timestamp
);

CREATE TABLE USER_GROUPS
(
    id         varchar(36) PRIMARY KEY,
    name       varchar(256) NOT NULL,
    deleted_at timestamp
);

CREATE TABLE MEMBERS
(
    id            varchar(36) PRIMARY KEY,
    user_group_id varchar(36)  NOT NULL,
    username      varchar(256) NOT NULL,
    deleted_at    timestamp,
    CONSTRAINT FK_user_group_id_member FOREIGN KEY (user_group_id) REFERENCES USER_GROUPS (ID)
);

CREATE TABLE USER_GROUP_WORKSPACES
(
    id            varchar(36) PRIMARY KEY,
    user_group_id varchar(36) NOT NULL,
    workspace_id  varchar(36) NOT NULL,
    permission    jsonb       NOT NULL,
    deleted_at    timestamp,
    CONSTRAINT fk_user_group_workspace_group FOREIGN KEY (user_group_id) REFERENCES USER_GROUPS (ID),
    CONSTRAINT fk_workspace_group_user_group FOREIGN KEY (workspace_id) REFERENCES WORKSPACES (ID)
);

CREATE TABLE DEPLOYMENTS
(
    id         varchar(36) PRIMARY KEY,
    name       varchar(255) NOT NULL,
    version    varchar(255) NOT NULL,
    circle_id  varchar(36) NOT NULL,
    deleted_at timestamp,
    CONSTRAINT fk_circle_deployment FOREIGN KEY (circle_id) REFERENCES CIRCLES (ID)
);

CREATE TABLE CIRCLE_USER_GROUPS
(
    id            varchar(36) PRIMARY KEY,
    user_group_id varchar(36) NOT NULL,
    circle_id     varchar(36) NOT NULL,
    deleted_at    timestamp,
    CONSTRAINT fk_user_group_circle FOREIGN KEY (user_group_id) REFERENCES USER_GROUPS (ID),
    CONSTRAINT fk_circle_user_group FOREIGN KEY (circle_id) REFERENCES CIRCLES (ID)
);

CREATE TABLE BUTLER_CONFIGURATIONS
(
    id           varchar(36) PRIMARY KEY,
    name         varchar(256) NOT NULL,
    url          varchar(256) NOT NULL,
    git_token    varchar(256) NOT NULL,
    workspace_id varchar(36)  NOT NULL,
    deleted_at   timestamp,
    CONSTRAINT fk_workspace_id_deployment_configuration FOREIGN KEY (workspace_id) REFERENCES WORKSPACES (ID)
);

CREATE TABLE DATASOURCES
(
    id            varchar(36) PRIMARY KEY,
    name          varchar(256) NOT NULL,
    type          varchar(256) NOT NULL,
    configuration jsonb        NOT NULL,
    workspace_id  varchar(36)  NOT NULL,
    deleted_at    timestamp,
    CONSTRAINT fk_workspace_id_datasource FOREIGN KEY (workspace_id) REFERENCES WORKSPACES (ID)
);

CREATE TABLE ACTIONS
(
    id            varchar(36) PRIMARY KEY,
    name          varchar(256) NOT NULL,
    type          varchar(256) NOT NULL,
    configuration jsonb        NOT NULL,
    description   varchar(256) NOT NULL,
    workspace_id  varchar(36)  NOT NULL,
    deleted_at    timestamp,
    CONSTRAINT fk_workspace_id_action FOREIGN KEY (workspace_id) REFERENCES WORKSPACES (ID)
);

