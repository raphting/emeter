Emeter
======

Convert dutch smart meter data to the Prometheus metrics format.

Use case
--------

I own a smart meter for electricity counting. To access the data, I ordered a
[cable](https://www.webshop.cedel.nl/Slimme-meter-kabel-P1-naar-USB) and attached it to the smart meter and my
Raspberry Pi 2.

My server is able to frequently gather the currently delivered kW. In that way I can display my energy consumption
in nearly realtime.

Configuration
-------------

* Port: 9688 hardcoded
* Path: /metrics
* Metric: `emeter_pwr_delivered` current energy consumption

Documentation
-------------

[The P1 standard](http://files.domoticaforum.eu/uploads/Smartmetering/DSMR%20v4.0%20final%20P1.pdf).
Unfortunately, I could not make the Checksum (CRC16) work. The documentation describes x^16 + x^15 + x^2 + 1 in LSB.
This converts to the so called _IBM Table_ with `0xA001`. For now I have to live with transmission glitches that are
really unlikely to happen.