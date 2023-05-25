# freertos_lwip_mbedtls7

Important reminders:

1. create project from `freertos_lwip4` template;
2. import mbedTLS sources manually;
3. replace `mbedtls_config.h` by the one from `mbedtls_get_cfg` project;
4. replace `net_sockets.h` by the one from eziya;
5. replace `lwipopts.h` by the one from `mbedtls_get_cfg` project;
6. import `hardware_rng.c` from `mbedtls_get_cfg` project.
