PGDMP     ,    
                {            postgres    15.2    15.2     �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    5    postgres    DATABASE        CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_Indonesia.1252';
    DROP DATABASE postgres;
                postgres    false            �           0    0    DATABASE postgres    COMMENT     N   COMMENT ON DATABASE postgres IS 'default administrative connection database';
                   postgres    false    3326                        3079    16384 	   adminpack 	   EXTENSION     A   CREATE EXTENSION IF NOT EXISTS adminpack WITH SCHEMA pg_catalog;
    DROP EXTENSION adminpack;
                   false                        0    0    EXTENSION adminpack    COMMENT     M   COMMENT ON EXTENSION adminpack IS 'administrative functions for PostgreSQL';
                        false    2            �            1259    16440    book    TABLE     �   CREATE TABLE public.book (
    id integer NOT NULL,
    title character varying(300) NOT NULL,
    author character varying(300) NOT NULL,
    deskripsi character varying(500) NOT NULL
);
    DROP TABLE public.book;
       public         heap    postgres    false            �            1259    16439    book_id_seq    SEQUENCE     �   CREATE SEQUENCE public.book_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 "   DROP SEQUENCE public.book_id_seq;
       public          postgres    false    216                       0    0    book_id_seq    SEQUENCE OWNED BY     ;   ALTER SEQUENCE public.book_id_seq OWNED BY public.book.id;
          public          postgres    false    215            f           2604    16443    book id    DEFAULT     b   ALTER TABLE ONLY public.book ALTER COLUMN id SET DEFAULT nextval('public.book_id_seq'::regclass);
 6   ALTER TABLE public.book ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    215    216    216            �          0    16440    book 
   TABLE DATA           <   COPY public.book (id, title, author, deskripsi) FROM stdin;
    public          postgres    false    216   �                  0    0    book_id_seq    SEQUENCE SET     9   SELECT pg_catalog.setval('public.book_id_seq', 9, true);
          public          postgres    false    215            h           2606    16447    book book_pkey 
   CONSTRAINT     L   ALTER TABLE ONLY public.book
    ADD CONSTRAINT book_pkey PRIMARY KEY (id);
 8   ALTER TABLE ONLY public.book DROP CONSTRAINT book_pkey;
       public            postgres    false    216            �   �   x�e�Kn� ���>��RU���E�q�� )��I�ݔH|���t3���V�T��*��5�����3�v���Y�?�}6a�!ZM��x��9n.�w��j��3|Q'\�t��P�Q�°P]�ىh�3��R
��#Iuu�,z�l	���A��P�z3C�uq�n�Еpm��DWO���g=\d����>,SH�Ξ�<@ۍ���?�9�a��=�m�j�S+R����:�<��i>B�#��     