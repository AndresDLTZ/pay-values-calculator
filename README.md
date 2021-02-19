<h1 align="center">Pay Values Calculator</h1>

## Table of Contents

- [Overview](#overview)
- [Approach and Methodology](#approach-and-methodology)
    - [Kanban](#kanban)
    - [Test-Driven Development TDD](#test-driven-development-tdd)
- [Architecture](#architecture)
- [Prerequisites](#prerequisites)
- [Build](#build)
- [Running unit tests](#running-unit-tests)
- [Licence](#licence)


## Overview

This is a project was implemented with [Go](https://golang.org/)
version [1.15.8](https://golang.org/dl/) to calculate values to be paid for hours worked
by employees based on the day of the week and time of day,
according to a specified rate.

Pay Values Calculator receives txt file path by keyboard input.
File has several data with the name of the employees, the day 
of the week and the time of the day worked.
File is read line by line, the pattern is validated, if it is
valid a goroutine is run to calculate the value to be paid 
(for each line of the file a goroutine is run).
The line is separated in name y day of the week with the time 
of the day worked. Each schedule is compared to the specified 
rate, and the value to be paid for each schedule is calculated, 
until the total value to be paid for the employee's hours worked
is obtained.

## Approach and Methodology

### Kanban

[Kanban](https://kanbanguides.org/) is an agile work management 
method created by David J. Anderson in Japan in the 1940s and 
used by Toyota engineers. It allows getting workflow information
in a single board, allows knowing which tasks are done, which 
tasks are delayed and allows estimating when the project will be
ready.

This project used the Kanban methodology, the Kanban board was 
divided into four groups "to do", "doing", "test" and "done".

The general tasks included in the Kanban board were: reading the file and its lines, 
searching in an array, operations with times, getting payment 
rate data, validations, pay value calculations and concurrent 
execution of the calculations.

### Test-Driven Development TDD
[TTD](http://agiledata.org/essays/tdd.html) is a software 
development approach where it writes the test first, then writing
and refactoring code to meet that test.

The project used TDD to think through its requirements and 
ensure they are met.

## Architecture

This project used design patterns to solve the requirements and 
provide a more robust and maintainable code.

The singleton creational design pattern is used to provide a 
single instance of the acme-daily-payment-rate.txt file to avoid 
having more than a single instance of the file and avoid resource
consumption. This was done in order to simulate that the daily 
payment rate data is get from a database.

The facade structural design pattern is used to provide the 
main's code with an interface to hide the complexity of the pay 
values calculations of each line of file and have clean code.

## Prerequisites

Download and install Go (https://golang.org/)

## Build

1. Run `go build`
  
    + Run `go build src/main.go`.
  
    or
    + Run `go build -o pay-values-calculator src/main.go`
      to set binary file name "pay-values-calculator".


2. It is necessary to create a json file in files directory
[files/acme-daily-payment-rate.json](https://github.com/AndresDLTZ/pay-values-calculator/blob/master/files/acme-daily-payment-rate.json)
together executable file.


3. Open command line terminal and execute `./main`.

```
Ingrese ruta del archivo con nombre y horarios (.txt): ./files/name-and-schedules1.txt

ASTRID [MO10:00-12:00 TH12:00-14:00 SU20:00-21:00]
El importe a pagar a ASTRID es: 85.00 USD

RENE [MO10:00-12:00 TU10:00-12:00 TH01:00-03:00 SA14:00-18:00 SU20:00-21:00]
El importe a pagar a RENE es: 215.00 USD

DAVID==MO10:00-12:00,TU10:00-12:00,TH01:00-03:00,SA14:00-18:00,SU20:00-21:00
Advertencia: nombre y horarios trabajados no se corresponde con el patrón

MONICA [MO10:00-09:00 TH12:00-14:00 SU20:00-21:00] 
Advertencia: No se puede calcular el horario a pagar, horario 10:00-09:00 fuera de referencia

ANNA [MO10:00-12:00 TH12:00-14:00 SU20:00-21:00]
El importe a pagar a ANNA es: 85.00 USD

JAQUE [MO10:00-12:00 TU10:00-12:00 TH01:00-03:00 SA14:00-18:00 SU20:00-21:00]
El importe a pagar a JAQUE es: 215.00 USD


Programa terminado :)
```

## Running unit tests

Unit test via Go [testing](https://golang.org/pkg/testing/) 
package. 

+ Run all test of package:
    + `go test ./src/pkg/<<package name>> -v`
    

+ Run specify test of package:
    + `go test ./src/pkg/<<package name>> -v -run <<test name>>`
  
## Licence

Andrés de la Torre
[MIT LICENCE](https://github.com/AndresDLTZ/pay-values-calculator/blob/master/LICENSE)