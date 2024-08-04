# Analizador de Texto Multifuncional en Go

¡Bienvenido al proyecto del Analizador de Texto Multifuncional en Go con bucles! Este repositorio acompaña al post "NIVEL 4: CREEMOS UN ANALIZADOR DE TEXTO MULTIFUNCIONAL EN GOLANG" de nuestro blog gognition.pro donde aprendimos a crear un analizador de texto usando bucles for.

## 📋 Descripción

Este proyecto implementa un analizador de texto multifuncional. El programa cuenta el número de palabras, identifica la palabra más larga, muestra la frecuencia de letras, invierte las palabras, busca palabras clave y analiza la estructura de oraciones.

## 🚀 Comenzando

Para utilizar este analizador de texto, asegúrate de tener Go instalado en tu sistema.

### Prerrequisitos

- Go (versión 1.15 o superior recomendada)

### Instalación

1. Haz un fork de este repositorio haciendo clic en el botón "Fork" en la parte superior derecha de esta página.

2. Clona tu fork a tu máquina local:

```bash
git clone https://github.com/TU_USUARIO/gognition-nivel5-analizador-textos-con-bucles.git
```

## 💻️ Uso

Para ejecutar el analizador de texto:
```bash
go run main.go
```

El programa mostrará el análisis del texto proporcionado, con salidas como:
```
Analizador de Texto Multifuncional
----------------------------------
1. Número total de palabras: 48
2. Palabra más larga: programación
3. Frecuencia de letras:
   Longitud de 'cafeteria': 9
   Longitud de 'cafetería': 10
   Longitud real de 'cafetería': 9
   f: 2
   g: 7
   s: 16
   a: 18
   ñ: 2
   o: 23
   c: 9
   m: 6
   i: 17
   h: 2
   á: 1
   u: 6
   í: 1
   ó: 2
   y: 4
   x: 2
   v: 1
   d: 10
   p: 8
   r: 16
   t: 9
   é: 2
   n: 16
   l: 10
   j: 2
   k: 2
   e: 30
   b: 4
   ú: 1
   Total de letras únicas: 29
4. Texto con palabras invertidas: oG se nu ejaugnel ed nóicamargorp etnerrucnoc y odalipmoc odaripsni ne al sixatnis ed .C aH odis odallorrased rop ,elgooG y sus serodañesid selaicini noreuf treboR ,remeseirG boR ekiP y neK .nospmohT oG se ,etneicife elbalacse y .ovitcudorp néibmaT ajenam otxet ne :loñapse ,á ,é ,í ,ó .ú
5. Número de palabras clave: 3
6. Análisis de estructura de oraciones:
   Oración 1: 15 palabras
   Oración 2: 17 palabras
   Oración 3: 6 palabras
   Oración 4: 10 palabras
```

## 🏗️ Compilación (Build)

Para compilar el proyecto y crear un ejecutable, puedes usar los siguientes comandos dependiendo de tu sistema operativo objetivo:
```bash
go build -o analizador-texto
```

Para Windows (desde Linux o macOS):
```bash
GOOS=windows GOARCH=amd64 go build -o analizador-texto.exe
```

Para macOS (desde Windows o Linux):
```bash
GOOS=darwin GOARCH=amd64 go build -o analizador-texto-mac
```

Para Linux (desde Windows o macOS):
```bash
GOOS=linux GOARCH=amd64 go build -o analizador-texto-linux
```

Después de la compilación, puedes ejecutar el programa usando:
```
./analizador-texto  # En Linux o macOS
analizador-texto.exe  # En Windows
```

## Visita gognition.pro https://www.gognition.pro/