# fan-control
CPU fan control with arduino

## Objetivo

Implementar uma controladora PWM para acionar os fans do servidor de forma a minimizar ruído mas mantendo a resposta em caso de demanda de CPU.

Será utilizado um Placa Minima Arduino Digispark Kickstarter Attiny85 Usb Dev

Também será desenvolvido um software em golang para fazer a interface entre os valores dos sensores de temperatura da CPU, cálculo da resposta para os fans e comunicação via porta serial.

## Versão 1

A controladora enviará periodicamente o valor do PWM atual, via serial, para controle do gerenciador.
O gerenciador fará o mapeamento das temperaturas e do PWM atual, e enviará via serial o valor do PWM caso seja necessário alterar.

## Versão 2

A controladora também monitorará a velocidade real dos fans e enviará a informação via serial para o gerenciador.


## Links

### Arduino

- https://forum.arduino.cc/t/12v-fan-speed-control-3-pins/671982/6
- https://www.irjmets.com/uploadedfiles/paper//issue_10_october_2023/45788/final/fin_irjmets1698948143.pdf
- https://forum.arduino.cc/t/how-do-i-control-a-12v-dc-fan-with-arduino/657589/3
- https://forum.arduino.cc/t/reading-fan-rpm-with-pwm/368626

### Gerenciador

- https://leapcell.io/blog/gopsutil-go-performance?ref=dailydev
- https://pkg.go.dev/go.bug.st/serial
