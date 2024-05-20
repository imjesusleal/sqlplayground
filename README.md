# sqlPlayground

Playground de sqlite3 in browser. Este proyecto se hizo en función de que cualquier persona que comience con alguna base de datos relacional, tenga la oportunidad de entrar y ejecutar consultas arbitrarias sin tener que crear todo el runtime en su sistema. 

Si bien es un proyecto simple, plantea una solución para aquellos que apenas estan entrando en el mundo del desarrollo (o de sistemas en general). En específico lo cree para mis compañeros del IFTS 18 que justo estamos cursando bases de datos y pense que les vendria bien. 

En sus proximas versiones estará desplegado y podrá ser accesible por cualquier persona, mientras tanto si deseas probarlo localmente puedes instalarlo siguiendo las instrucciones que te dejare abajo. 

# Instalación local

Para poder recrearlo dentro de tu red local puedes simplemente servir los archivos de la carpeta static. Para servir los archivos puedes utilizar cualquier servidor http que prefieras, algo muy común es usar python para esto:

```
python -m http.server {port}
```

Donde port simplemente es el puerto que quieras utilizar para servir los archivos (generalmente 80, 8080 u 3000). Pero si ya tienes Go instalado puedes simplemente utilizar la regla del Makefile para servir los archivos:

```
make serve
```

Ya solo quedaría probarlo y decirme que les parece. 
