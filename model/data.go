package model

import "strings"

var alphabet = []string{"e", "t", "o", "s", "a", "r", "n", "i", "d", "l", "c", "u", "m", "p", "b", "g", "v", "y", "q", "h", "f", "z", "j", "ñ", "x", "k", "w"}

func GetSuggestion(exclusions string) string {
	for _, value := range alphabet {
		if !strings.Contains(exclusions, value) {
			return value
		}
	}
	return ""
}

/*
Frecuencia de Aparición de letras en español
https://es.wikipedia.org/wiki/Frecuencia_de_aparici%C3%B3n_de_letras
e	13,68%
t	12,63%
o	8,68%
s	7,98%
a	7,53%
r	6,87%
n	6,71%
i	6,25%
d	5,86%
l	4,97%
c	4,68%
u	3,93%
m	3,15%
p	2,51%
b	1,42%
g	1,01%
v	0,90%
y	0,90%
q	0,88%
h	0,70%
f	0,69%
z	0,52%
j	0,44%
ñ	0,31%
x	0,22%
k	0,02%
w	0,01%
*/
