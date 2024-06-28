# sqlPlayground

Playground de sqlite3 in browser. Este proyecto se hizo en función de que cualquier persona que comience con alguna base de datos relacional, tenga la oportunidad de entrar y ejecutar consultas arbitrarias sin tener que crear todo el runtime en su sistema. 

Si bien es un proyecto simple, plantea una solución para aquellos que apenas estan entrando en el mundo del desarrollo (o de sistemas en general). En específico lo cree para mis compañeros del IFTS 18 que justo estamos cursando bases de datos y pense que les vendria bien. 

En sus proximas versiones estará desplegado y podrá ser accesible por cualquier persona, mientras tanto si deseas probarlo localmente puedes instalarlo siguiendo las instrucciones que te dejare abajo. 

# Instalación local

Necesitas Golang para poder compilar el codigo hacia arquitectura WASM.

La instalación de las dependencias las deje todas en el Makefile, para correlo simplemente ejecuta:

```
make all
```

Con esto ya deberias tener todas las dependencias que necesitarias para poder contribuir en el proyecto.

# Agradecimientos

Todo el CSS ha sido generado utilizando Tailwind y, en particular, los componentes de la libreria de DaisyUI, pueden chequear su web: https://daisyui.com/
