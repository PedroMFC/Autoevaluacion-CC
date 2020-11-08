# 4. Examinar la estructura de capas que se forma al crear imágenes nuevas a partir de contenedores que se hayan estado ejecutando.

Para examinar las capas de la imagen anterior ejecutamos el comando 

```
sudo jq '.' /var/lib/docker/image/overlay2/imagedb/content/sha256/0ab7a52240de9bd2d9531e2ba4a5e2c8c1bb5973f136fa5e667af638dfcd1efb

```

Si vamos al elemento `diff_ids` nos muestra

![](./imgs/4.png)