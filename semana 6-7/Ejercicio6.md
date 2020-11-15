# 6. Reproducir los contenedores creados anteriormente usando un `Dockerfile`.

Se ha copiado en el `Dockerfile` el código proporcionado. Ahora creamos la imagen mediante `build` y vemos que efectivamente se ha creado.

![](./imgs/6.1.png)

![](./imgs/6.2.png)

Ahora mediante la orden `inspect`, vemos que efectivamente esta imagen contiene muchas capas.

![](./imgs/6.3.png)

Cambiamos el contenido del [Dockerfile](./Dockerfile) por el nuevo. Veamos el resultado.

![](./imgs/6.4.png)

![](./imgs/6.5.png)

Vemos que efectivamente ha reducido su tamaño. Comprobemos las capas.

![](./imgs/6.6.png)