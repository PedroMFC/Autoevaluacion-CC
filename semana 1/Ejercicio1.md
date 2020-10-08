# Buscar una aplicación de ejemplo, preferiblemente propia, y deducir qué patrón es el que usa. ¿Qué habría que hacer para evolucionar a un patrón tipo microservicios?

En la asignatura PDOO del grado en informática, se desarrolló una apliacación llamada Napakalaki que consistía en un juego de cartas. La arquitectura que seguía esa aplicación es el conocido como Modelo-Vista-Controlador.

Para evolucionarlo a una arquitectura de microservicios se tendrían que dividir las funcionalidades que tiene la aplicación. Así, habría que hacer un microsrevicio para manejar las cartas disponibles, otro para llevar la lógica de otro, otro para interactuar con los jugadores, etc.

