if theres a patent for regex manipulation of gcode or inverting the axis in a file or something im gonna be pissed\
with that sed(1) i needed a script to invert my z axis in gcode files before running them (my z axis is somehow reversed and i didnt feel like messing with grbl)\
so heres some go code to do that\
fair warning it can get a -0.0000 so if your machine controller doesnt like that... its a 100< line script figure it out
