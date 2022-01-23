# mlsmpm-particles-go

This repository simulates and renders fluid particles in two dimensions.  
[Click HERE to run the simulation in a browser](https://robkau.github.io/mlsmpm-particles-go/).  

The simulation uses MLS-MPM algorithm (Moving Least Squares Material Point Method).  
I implemented this by following [nialltl's article on MLS-MPM](https://nialltl.neocities.org/articles/mpm_guide.html) and [matching example code](https://github.com/nialltl/incremental_mpm).

Library [ebiten](https://github.com/hajimehoshi/ebiten) is used to render the output to a window and allows seamless compilation to WASM.   
The simulation is single threaded but still renders the examples in real time on a fast CPU.  

Running this in WASM in the browser is about 10x slower than native speed. WebGL support may help when it finally arrives...

Build and run the _cmd/sim_ package to run simulation at native speed.


---

## Example Videos



https://user-images.githubusercontent.com/1654124/150659593-25e9022c-a27d-441b-9481-8f8f748cbf85.mov



https://user-images.githubusercontent.com/1654124/150659675-e4b4bea0-cd13-49f4-a51f-5066abe4db6e.mov



https://user-images.githubusercontent.com/1654124/150659679-2d048d5e-98a0-4d85-ad9f-ac23a4e63c24.mov



https://user-images.githubusercontent.com/1654124/150659685-a4d55341-2a1d-4031-870a-860f0e82c444.mov



https://user-images.githubusercontent.com/1654124/150659686-dc8009ae-6af3-44e0-8af6-0ed701736cb7.mov



https://user-images.githubusercontent.com/1654124/150659601-aa409bea-557c-44e2-ae13-d61cc1c9d609.mov

