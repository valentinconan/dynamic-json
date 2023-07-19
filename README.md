# dynamic-json

## Main purpose
This library allow to unmarshal json data into a declared structure, as you can do using "encoding/json" library but add the possibility to manage unknown fields in a map.

This will allow you to keep these data even if you don't need then 

Case study: an api application that intercepts requests and manipulates known data, but needs all the data, known and unknown, in order to transmit the original request to the main target.
 

## Technical details
The library use `encoding/json` libraries

