# mlsmpm-particles-go

This repository simulates and renders fluid particles in two dimensions.   
The simulation uses MLS-MPM algorithm (Moving Least Squares Material Point Method).  
I implemented this by following the examples from [this excellent webpage](https://nialltl.neocities.org/articles/mpm_guide.html) and [example code](https://github.com/nialltl/incremental_mpm).

Library [ebiten](https://github.com/hajimehoshi/ebiten/issues) is used to render the output to a window.  
The simulation is single threaded but still renders the examples in real time on a fast CPU.

Build and run the _cmd/sim_ package to interact with realtime simulation.

---

## Examples

![Falling squarelets 1](renders/i-0-30.mkv?raw=true "Falling squarelets 1")

![Falling squarelets 2](renders/i-60-30.mkv?raw=true "Falling squarelets 2")

![Streamers](renders/output-s-rendered512-144-interp.mkv "Particle streamers")