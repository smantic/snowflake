# smantic/snowflake 

I thought it was weird that existing implementations of snowflake ID didn't utilize all 64 bits 
So I wrote a library that implements a signed timestamp snowflake ID.

![format](https://upload.wikimedia.org/wikipedia/commons/5/5a/Snowflake-identifier.png)

There are only 63 bits used in this reference, the missing bit is the most significant bit, which is the sign bit. 

We think of the timestamps as the number of milliseconds since a chosen epoch, so generally the timestamps in snowflake IDs are unsigned.
The problem arises when we consider Some languages ( *JavaScript* ) which don't have unsigned int types. If we have a unsigned timestamp that uses the MSB, all the sudden our 
snowflake ID is a negative number, which breaks our sorting.

for an epoch of 2015

-Something < something < something 
2090 <  2020 < 2021 

One thing you could do is just start using signed timestamps and set the epoch way in the future. 
So I decided to implement a library showing this for fun.


### Spaces

42 bit timestamp = ~140 years
41 Bit timestamp = ~70 Years 

We lose 1/2 of our time when using 41 bits. Twitter and discord probably wont be around in 70 years, so it's likely that it wont matter. To put that in perspective Turing created the Turing Machine 85 years ago.


## Signed timestamps

Signed timestamps are actually a thing but you may not have seen or used them. A negative timestamp indicates a time before your selected epoch.

To take full advantage of all 42 bits on a signed timestamp, we can set the epoch to be 69 years in advance. 


## Other options 

We can also use a signed timestamp, but allocate a bit else where in the snowflake structure. 
For example we could use 13 bits instead of 12 for the sequence, allowing for 4096 (2x) more unique IDs per timestamp per machine/instance. 
