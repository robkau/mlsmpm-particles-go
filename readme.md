# mlsmpm-particles-go

This repository simulates and renders fluid particles in two dimensions.   
The simulation uses MLS-MPM algorithm (Moving Least Squares Material Point Method).  
I implemented this by following the examples from [this excellent webpage](https://nialltl.neocities.org/articles/mpm_guide.html) and [example code](https://github.com/nialltl/incremental_mpm).

Library [ebiten](https://github.com/hajimehoshi/ebiten/issues) is used to render the output to a window.  
The simulation is single threaded but still renders the examples in real time on a fast CPU.

Build and run the _cmd/sim_ package to interact with realtime simulation.

---

## Example Videos

![Falling squarelets 1](renders/i-0-30.mp4?raw=true "Falling squarelets 1")

[Falling squarelets 2](renders/i-60-30.mkv?raw=true "Falling squarelets 2")

[Falling squarelets 3](renders/i-120-30.mkv?raw=true "Falling squarelets 3")

[Falling squarelets 4](renders/i-180-30.mkv?raw=true "Falling squarelets 4")

[Falling squarelets 5](renders/i-210-30.mkv?raw=true "Falling squarelets 5")

[Streamers](renders/output-s-rendered512-144-interp.mkv "Particle streamers")