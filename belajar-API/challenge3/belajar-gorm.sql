PGDMP         %                {            belajar-gorm    15.2    15.2                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16470    belajar-gorm    DATABASE     �   CREATE DATABASE "belajar-gorm" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
    DROP DATABASE "belajar-gorm";
                postgres    false            �            1259    16471    buku    TABLE     E  CREATE TABLE public.buku (
    id bigint NOT NULL,
    judul_buku character varying(200) NOT NULL,
    author character varying(200) NOT NULL,
    kategori character varying(200) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);
    DROP TABLE public.buku;
       public         heap    postgres    false            �            1259    16481    bukus    TABLE     F  CREATE TABLE public.bukus (
    id bigint NOT NULL,
    judul_buku character varying(200) NOT NULL,
    author character varying(200) NOT NULL,
    kategori character varying(200) NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);
    DROP TABLE public.bukus;
       public         heap    postgres    false            �            1259    16480    bukus_id_seq    SEQUENCE     u   CREATE SEQUENCE public.bukus_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.bukus_id_seq;
       public          postgres    false    216            	           0    0    bukus_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.bukus_id_seq OWNED BY public.bukus.id;
          public          postgres    false    215            k           2604    16484    bukus id    DEFAULT     d   ALTER TABLE ONLY public.bukus ALTER COLUMN id SET DEFAULT nextval('public.bukus_id_seq'::regclass);
 7   ALTER TABLE public.bukus ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    215    216    216                       0    16471    buku 
   TABLE DATA           X   COPY public.buku (id, judul_buku, author, kategori, created_at, updated_at) FROM stdin;
    public          postgres    false    214   
                 0    16481    bukus 
   TABLE DATA           Y   COPY public.bukus (id, judul_buku, author, kategori, created_at, updated_at) FROM stdin;
    public          postgres    false    216   '       
           0    0    bukus_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.bukus_id_seq', 6, true);
          public          postgres    false    215            o           2606    16479    buku buku_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY public.buku
    ADD CONSTRAINT buku_pkey PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.buku DROP CONSTRAINT buku_pkey;
       public            postgres    false    214            q           2606    16490    bukus bukus_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.bukus
    ADD CONSTRAINT bukus_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.bukus DROP CONSTRAINT bukus_pkey;
       public            postgres    false    216                   x������ � �            x�}�ˎ� E��ޏ�ɷt�&��4$�1�?��Y���X���1�d��(�2�Zϔ��xԿL	��� �� ���$"~����6ҍ(��VJ%L�hG����z��B1��i%*��U/�#k�T�p�POp�q&��2EFa���T<%��V{�S���z��e��k�o߲>8���e�%�=�ıQҴm���L�f.���6N�o�Х�\����R����,�A�ji�۲p��F;s4�1�S�9����k     