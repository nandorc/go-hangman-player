# Hangman Player

> **Team name:** #radiopragma
> 
> **Unique team number:** 2
> 
> **Team members:**
> 
> - Carlos Alberto Lopez Mazo <carlos.lopez@pragma.com.co>
> - John David Gonzalez Alzate <john.gonzalez@pragma.com.co>
> - Daniel Fernando Rivera Cubillos <daniel.rivera@pragma.com.co>

## Description

This project serves an API which provides a player for the classic game *hangman*. It's made to be served through the port 8090 on localhost (http://localhost:8090)

## Endpoints

> `POST /hangman`
> 
> Entry point to receive character suggestions bases on a word to guess and a regsiter of used characters.
> 
> **payload**
> >     {
> >         "palabra" : string
> >         "intentos" : string
> >     }
> 
> **response**
> >     {
> >         "letra" : string
> >     }