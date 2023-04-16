
-- auto update update_time when update
CREATE OR REPLACE FUNCTION update_modified_column()   
RETURNS TRIGGER AS $$
BEGIN
    NEW.update_time = now();
    RETURN NEW;   
END;
$$ language 'plpgsql';


--- user tablse

CREATE TYPE user_role AS ENUM('admin', 'general');
CREATE TYPE public.gender AS ENUM ('unknown', 'female', 'male');
COMMENT ON TYPE public.gender IS '性别';

-- Table: public.user

-- DROP TABLE IF EXISTS public."user";

CREATE TABLE public."user" (
	id bigserial NOT NULL, -- 用户id
	username varchar(16) NOT NULL,
	email varchar(64) NULL, -- 邮箱
	nickname varchar NOT NULL,
	"role" public.user_role NOT NULL DEFAULT 'general'::user_role, -- 用户角色
	gender public.gender NOT NULL DEFAULT 'unknown'::gender, -- 性別
	"password" varchar(128) NULL,
	create_time timestamptz(0) NOT NULL DEFAULT now(), -- 创建时间
	update_time timestamptz(0) NOT NULL DEFAULT now(), -- 修改时间        
	CONSTRAINT user_name_unique_idx UNIQUE (username),
	CONSTRAINT user_pk PRIMARY KEY (id)
);
COMMENT ON TABLE public."user" IS '用户表，包含管理用户和普通用户';

-- Column comments

COMMENT ON COLUMN public."user".id IS '用户id';
COMMENT ON COLUMN public."user".username IS '用户登录的用户名';
COMMENT ON COLUMN public."user".email IS '邮箱';
COMMENT ON COLUMN public."user".nickname IS '用于显视的用户昵称';
COMMENT ON COLUMN public."user"."role" IS '用户角色';
COMMENT ON COLUMN public."user".gender IS '性別';
COMMENT ON COLUMN public."user".create_time IS '创建时间';
COMMENT ON COLUMN public."user".update_time IS '修改时间';

-- Constraint comments

COMMENT ON CONSTRAINT user_name_unique_idx ON public."user" IS 'username needs unique';

-- Table Triggers

create trigger update_user_update_time before
update
    on
    public."user" for each row execute function update_modified_column();


-- post  tables

CREATE TYPE public.post_status AS ENUM
    ('draft', 'private', 'published');
COMMENT ON TYPE public.post_status
    IS '文章状态，  draft: 草稿, private 仅自己可见, published 发布状态';


-- public.post definition

-- Drop table

-- DROP TABLE public.post;

CREATE TABLE public.post (
	id bigserial NOT NULL,
	title varchar(256) NOT NULL, -- 文章标题
	body text NULL, -- 文章正文
	status public.post_status NOT NULL DEFAULT 'draft'::post_status, -- 文章状态，draft: 草稿, private 仅自己可见, published 发布状态
	category_id int8 NULL, -- 文章分类的id 关联category表
	user_id int8 NOT NULL, -- 文章作者id 对应user表中的 id
	create_time timestamptz NOT NULL DEFAULT now(), -- 文章创建时间
	update_time timestamptz NOT NULL DEFAULT now(), -- 文章修改时间
	textsearch tsvector NULL GENERATED ALWAYS AS (to_tsvector('jiebaqry'::regconfig, (COALESCE(title, ''::character varying)::text || ' '::text) || COALESCE(body, ''::text))) STORED,
	CONSTRAINT post_body_unique_idx UNIQUE (title),
	CONSTRAINT post_pkey PRIMARY KEY (id)
);
CREATE INDEX post_category_id_index ON public.post USING btree (category_id);
COMMENT ON INDEX public.post_category_id_index IS '文章分类索引';
CREATE INDEX post_status_index ON public.post USING btree (status);
COMMENT ON INDEX public.post_status_index IS '文章状态索引';
CREATE INDEX textsearch_idx ON public.post USING gin (textsearch);
COMMENT ON TABLE public.post IS '文章表';

-- Column comments

COMMENT ON COLUMN public.post.title IS '文章标题';
COMMENT ON COLUMN public.post.body IS '文章正文';
COMMENT ON COLUMN public.post.status IS '文章状态，draft: 草稿, private 仅自己可见, published 发布状态';
COMMENT ON COLUMN public.post.category_id IS '文章分类的id 关联category表';
COMMENT ON COLUMN public.post.user_id IS '文章作者id 对应user表中的 id';
COMMENT ON COLUMN public.post.create_time IS '文章创建时间';
COMMENT ON COLUMN public.post.update_time IS '文章修改时间';

-- Constraint comments

COMMENT ON CONSTRAINT post_body_unique_idx ON public.post IS '文章标题唯一索引';

-- Table Triggers

create trigger update_user_update_time before
update
    on
    public.post for each row execute function update_modified_column();


--- 分类表
-- public.category definition

-- Drop table

-- DROP TABLE public.category;

CREATE TABLE public.category (
	id bigserial NOT NULL,
	"name" varchar(128) NOT NULL,
	user_id int8 NOT NULL,
	create_time timestamptz NOT NULL DEFAULT now(),
	update_time timestamptz NULL DEFAULT now(),
	CONSTRAINT category_pkey PRIMARY KEY (id)
);
CREATE INDEX category_user_id_index ON public.category USING btree (user_id);
COMMENT ON INDEX public.category_user_id_index IS 'category表user_id索引， 业务场景中，都查某个用户下的分类';
COMMENT ON TABLE public.category IS '文章分类表';

-- Column comments
COMMENT ON COLUMN public.category.id IS '分类的id,在post表中引用';
COMMENT ON COLUMN public.category."name" IS '分类名称，用户级別可以重复，重复了就复显视';
COMMENT ON COLUMN public.category.user_id IS '分类所属的用户，每个用户只能用自己的分类';
COMMENT ON COLUMN public.category.create_time IS '分类创建时间';
COMMENT ON COLUMN public.category.update_time IS '分类修改时间';

-- Table Triggers

create trigger update_category_update_time before
update
    on
    public.category for each row execute function update_modified_column();

COMMENT ON TRIGGER update_category_update_time ON public.category IS '当更新数时，自动更新update_time';


--- 设置表
-- public.setting definition

-- Drop table

-- DROP TABLE public.setting;

CREATE TABLE public.setting (       
        "key" varchar(128) NOT NULL,
        "data" jsonb NOT NULL,
	create_time timestamptz NOT NULL DEFAULT now(),
	update_time timestamptz NULL DEFAULT now(),
	CONSTRAINT setting_pkey PRIMARY KEY ("key")
);
COMMENT ON TABLE public.setting IS '设置表，包括blog2系统的一切设置, 每一条记录为一项设置，这样设计方便扩展。';

-- Column comments
COMMENT ON COLUMN public.setting."key" IS '设置的key,全局唯一,由业务中自己定义';
COMMENT ON COLUMN public.setting."data" IS '设置的内容，由于每项设置的内容不一样，因此为jsonb，兼容性强';
COMMENT ON COLUMN public.setting.create_time IS '设置创建时间';
COMMENT ON COLUMN public.setting.update_time IS '设置修改时间';


-- Table Triggers

create trigger update_setting_update_time before
update
    on
    public.setting for each row execute function update_modified_column();

COMMENT ON TRIGGER update_setting_update_time ON public.setting IS '当更新数时，自动更新update_time';
