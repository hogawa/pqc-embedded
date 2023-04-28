# lwip_bare

A simple test of Ethernet and LWIP TCP/IP stack configuration for the Nucleo-F207ZG board. Adapted from eziya's blog, where a Nucleo-F429ZI board is used:

> https://blog.naver.com/eziya76/221852430347

### Notes

When enabling the LWIP stack in CubeIDE, we have to enable the `LWIP_NETIF_LINK_CALLBACK` option under 'LWIP > Configurations > Key Options', as described by [eziya](https://blog.naver.com/eziya76/221852430347). In the current CubeIDE 1.12.0, the Ethernet clock is automatically configured with the option of "default board peripherals", so there's no need to manually configure it.

### Simple TCP client

On top of Ethernet initialization check, a simple TCP client:

> https://blog.naver.com/PostView.naver?blogId=eziya76&logNo=221862499239&parentCategoryNo=&categoryNo=38&viewDate=&isShowPopularPosts=false&from=postView

This is not yet fully functional, as the packet format needs to be adjusted on the server side. At least, we were able to detect returning packets from the server.
