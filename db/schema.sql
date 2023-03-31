create table results
(
    id         int                                not null auto_increment,
    x          int                                not null,
    y          int                                not null,
    sum        int                                not null,
    created_at datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP,
    constraint message_pk
        primary key (id),
    constraint message_pk
        unique (id)
);
