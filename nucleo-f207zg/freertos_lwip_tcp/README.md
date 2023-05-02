# freertos_lwip_tcp

FreeRTOS and LWIP running on the Nucleo-F207ZG board. This application sends GET requests to a server that responds with timestamp strings. It is based on eziya's example project:

> https://blog.naver.com/PostView.naver?blogId=eziya76&logNo=221871231508&parentCategoryNo=&categoryNo=38&viewDate=&isShowPopularPosts=false&from=postView

The sample server application is located in the `${PROJ_ROOT}/util/go_tstamp_srv/` folder, and is implemented using Golang (special thanks to @williamszk).

### Important configurations

When configuring FreeRTOS and LWIP on CubeIDE:

1. System Core > SYS > Timebase Source = TIM6
2. FreeRTOS > Interface = CMSIS_V1
3. FreeRTOS > Config Parameters > MINIMAL_STACK_SIZE = 256 words (as suggested by [eziya](https://blog.naver.com/PostView.naver?blogId=eziya76&logNo=221867311729&parentCategoryNo=&categoryNo=38&viewDate=&isShowPopularPosts=false&from=postView))
4. FreeRTOS > Config Parameters > TOTAL_HEAP_SIZE = 32768 Bytes (as suggested by [eziya](https://blog.naver.com/PostView.naver?blogId=eziya76&logNo=221867311729&parentCategoryNo=&categoryNo=38&viewDate=&isShowPopularPosts=false&from=postView))
5. FreeRTOS > Advanced settings > USE_NEWLIB_REENTRANT = Enabled
6. LWIP > RTOS_USE_NEWLIB_REENTRANT = 100 (as suggested in this ST forum thread [link](https://community.st.com/s/question/0D53W00002EBsjUSAT/stm32f207-lwip-freertos-configuration-error-rtosusenewlibreentrant))
