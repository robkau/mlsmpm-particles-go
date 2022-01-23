# mlsmpm-particles-go

This repository simulates and renders fluid particles in two dimensions.   
The simulation uses MLS-MPM algorithm (Moving Least Squares Material Point Method).  
I implemented this by following the examples from [this excellent webpage](https://nialltl.neocities.org/articles/mpm_guide.html) and [example code](https://github.com/nialltl/incremental_mpm).

Library [ebiten](https://github.com/hajimehoshi/ebiten/issues) is used to render the output to a window.  
The simulation is single threaded but still renders the examples in real time on a fast CPU.

Build and run the _cmd/sim_ package to interact with realtime simulation.

---

## Example Videos



https://user-images.githubusercontent.com/1654124/150659593-25e9022c-a27d-441b-9481-8f8f748cbf85.mov



https://user-images.githubusercontent.com/1654124/150659675-e4b4bea0-cd13-49f4-a51f-5066abe4db6e.mov



https://user-images.githubusercontent.com/1654124/150659679-2d048d5e-98a0-4d85-ad9f-ac23a4e63c24.mov



https://user-images.githubusercontent.com/1654124/150659685-a4d55341-2a1d-4031-870a-860f0e82c444.mov



https://user-images.githubusercontent.com/1654124/150659686-dc8009ae-6af3-44e0-8af6-0ed701736cb7.mov



https://user-images.githubusercontent.com/1654124/150659601-aa409bea-557c-44e2-ae13-d61cc1c9d609.mov

