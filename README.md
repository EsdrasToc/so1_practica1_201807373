#Practica 1
- Esdras Rodolfo Toc Hi
- 201807373


##Frontend

Para realizar el frontend de la aplicacion se utilizo el framework React y para el estilo del mismo se utilizo bootstrap.

-**Version de node:** 16.16.0
-**Version de npm:** 8.17.0

##Backend

Para el desarrollo del backend se utilizo el lenguaje Go, adicional a esto utilizamos el paquete gorilla/mux con la finalidad de levantar una rest API que servira como intermediario entre el frontend y la base de datos.

-**Version de Go:** 1.13.8
-**Paquete gorilla/mux:** https://github.com/gorilla/mux

##Base de datos

Al ser datos bastante simples de almacenar, se tomo la decision de utilizar un lenguage NoSQL, para ser mas exactos, utilizamos MongoDB.

Al momento de levantar los contenedores del programa, se creara una carpeta extra al lado de nuestras carpetas del frontend y backend llamada Data, es muy importante NO eliminar esta carpeta ya que en la misma se guardaran todas los datos almacendos por MongoDB.

##Instalacion

1. Instalar docker
	**Version de docker:** 20.10.12
2. Instalar docker compose
	**Version de docker-compose:** 1.29.2
3. Descargar el codigo fuente
4. El unico cambio que haremos en el codigo fuente para su funcionamiento sera dentro del frontend, ya que en los links utilizados para comunicarnos con el backend deberemos cambiar la IP por la ip de nuestro servidor. Esta es la unica modificacion que debera hacer a su codigo.
5. Utilizar el comando: docker-compose up

**Felicidades! has instalado con exito la aplicacion**
