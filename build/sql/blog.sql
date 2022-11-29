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
COMMENT ON COLUMN public."user".email IS '邮箱';
COMMENT ON COLUMN public."user".create_time IS '创建时间';
COMMENT ON COLUMN public."user".update_time IS '修改时间';
COMMENT ON COLUMN public."user"."role" IS '用户角色';
COMMENT ON COLUMN public."user".gender IS '性別';

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
	status public.post_status NOT NULL DEFAULT 'draft'::post_status, -- '文章状态，  draft: 草稿, private 仅自己可见, published 发布状态';
	tag_ids _int8 NULL, -- 文章标签，搜索使用
	user_id int8 NOT NULL, -- 文章作者id 对应user表中的 id
	create_time timestamptz NOT NULL DEFAULT now(), -- 文章创建时间
	update_time timestamptz NOT NULL DEFAULT now(), -- 文章修改时间
	CONSTRAINT post_body_unique_idx UNIQUE (title),
	CONSTRAINT post_pkey PRIMARY KEY (id)
);
COMMENT ON TABLE public.post IS '文章表';

-- Column comments

COMMENT ON COLUMN public.post.title IS '文章标题';
COMMENT ON COLUMN public.post.body IS '文章正文';
COMMENT ON COLUMN public.post.status IS '''文章状态，  draft: 草稿, private 仅自己可见, published 发布状态'';';
COMMENT ON COLUMN public.post.tag_ids IS '文章标签，搜索使用';
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
