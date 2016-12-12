#------------------
# Introduction
#------------------
This is a basic API code implemented using Go.
This API makes searchable information from a SQLite database and 
deliver the same in JSON format.

#------------------
# BVC Description
#------------------
This API implemented in Go, allows to deliver data stored on a 
SQLite database in JSON format.
To the API implementation, it was used a Go API Framework called Goji.
Goji is a minimalistic and flexible HTTP request multiplexer for 
Go (golang) https://goji.io

#------------------
# Requirements
#------------------
This program requires the following modules:
 
    * Go version 1.7.4 (go1.7.4.windows-amd64)
    * GCC (tdm64-gcc-5.1.0-2)
    * Goji
    * go-sqlite3
    * Glide version 0.12.3

#---------------------
# Recommended modules
#---------------------
    * SQLite connections 
        https://github.com/mattn/go-sqlite3
    
    * Glide
        https://github.com/Masterminds/glide/releases

    * GCC
        tdm64-gcc-5.1.0-2

#------------------
# Configuration
#------------------
It is necessary to set the environment variables after Go installation.
For more information about that, check the following link [Instalación de GO (Golang)](https://medium.com/@golang_es/instalaci%C3%B3n-de-go-golang-6fd5d7b9eb48#.u7ap68jv1)

Also make sure that both gcc compiler and glide environment variables are set

#------------------
# Troubleshooting
#------------------
If you have any trouble during the installation or testing check the following
statements

    * Context package error during Goji installation: 
        This goji error was solved installing a go version after 1.6.4
        Link to download go (https://golang.org/dl/)
    
    * exec: "gcc": executable file not found in %PATH%
        This error is solved installing tdm64-gcc-5.1.0-2.exe and removing
        another previous installation of gcc or installations related with mingw

    * cc1.exe: sorry, unimplemented: 64-bit mode not compiled in
        If it error appears, make sure you have a 64 bit installation 
        of tdm64-gcc-5.1.0-2 or another lib installation.


# -------------------------------------------
# Here are some routes to test the program
# -------------------------------------------

## General Search (filters by song, artist and genre)
localhost:8000/generalsearch/Rock

## Search by Song 
localhost:8000/searchbysong/Jude

## Search by Artist
localhost:8000/searchbyartist/Santana

## Search by Genre
localhost:8000/searchbygenre/Rock

## Songs By Length
localhost:8000/songsbylength/min=200&max=240

## List of genres
localhost:8000/genreslist
