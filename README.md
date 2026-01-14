## TCC: Estudando Autômatos Celulares com a Linguagem de Programação Go
## TCC: Studying Cellular Automata Using the Go Programming Language

_Lenia code has been updated. (2026/01)_

#### Studying Game of Life, SmoothLife, and Lenia

This repository contains the source code, simulations, and written project (TCC) for a study on Cellular Automata (CA) implemented in the `Go` programming language. The project explores the evolution from discrete models (Conway’s Game of Life) to continuous models (SmoothLife and Lenia).

### Basic Tutorial
**Requirements:** `Go`, `Gnuplot` and `FFmpeg`.

**Run a simulation:**

```bash
go run simulation-name.go
```
Each simulation generates `.dat` files representing the grid state for each generation (`NG`).

**Generate visuals:** Make changes to the provided `.plt` scripts (if necessary) and run in `Gnuplot` to convert `.dat` outputs into visuals:
- Use `grid.plt` to generate `.png` images of the grid.
- Use `pop-density.plt` to generate `.pdf` population density curves.

**Create video:** 

```bash
ffmpeg -r 7 -i image-%d.png -c:v libx264 -pix_fmt yuv420p video.mp4
```

### Features & Implementation
- All simulations are implemented in pure `Go` (no external dependencies) and utilize periodic boundary conditions.
- **Game of Life:** Discrete 2D implementation based on Moore neighborhoods.
- **SmoothLife:** Transition to continuous space/time following Rafler’s paper.
- **Lenia:** Continuous automata utilizing a manual 2D convolution kernel as described in Chan’s research.
- **Initial Conditions:** There are two initial conditions, one used in `GoL` and one used in `SmoothLife` and `Lenia`.
- **Data Logging:** Integrated `ic` (initial condition) and `op` (file operations) functions for automated data recording.

(**Note on Lenia:** While species form successfully, they may currently be unstable over long generations; updates are applied as optimizations continue.)

### Repository Contents
- **Code:** Updated (2025/2026) `.go` source files.
- **Videos:** `.mp4` results of simulations using specific initial conditions.
- **Plots:** `.plt` scripts for `Gnuplot` visualization.
- **Documents:** `.pdf` files (PT-BR).

### References & Acknowledgements
#### Acknowledgements
Special thanks to:
- **Supervision:** Professor Breno Ferraz de Oliveira.
- **Support:** Luiz Felipe Locatelli Giroldo.

#### Primary Resources
- [May-Leonard RPS Model `.pdf`](Docs/May_Leonard_RPS.pdf) & [`.go`](Code/ml_rps.go)

#### SmoothLife
- [duckythescientist's Python implementation](https://github.com/duckythescientist/SmoothLife/)
- [ionreq's application](https://sourceforge.net/projects/smoothlife/) & [YouTube channel](https://www.youtube.com/channel/UC_xsxCHaz_h-GGtOaFRGjvg)
- [0fps.net](https://0fps.net/tag/smoothlife/)
- [Rudyon's raylib-go code](https://github.com/rudyon/smoothlife-go/)
- [tsoding's C implementation](https://github.com/tsoding/SmoothLife/tree/master)

#### Lenia
- [Bert Wang-chak Chan (Chakazul) - Lenia portal](https://chakazul.github.io/lenia.html)
- [Coding Tutorial (Colab)](https://colab.research.google.com/github/OpenLenia/Lenia-Tutorial/blob/main/Tutorial_From_Conway_to_Lenia.ipynb)
- [Source Repository](https://github.com/Chakazul/Lenia)
- [Online Lenia Simulator](https://chakazul.github.io/Lenia/JavaScript/Lenia.html) & [Primordia Simulator](https://chakazul.github.io/Primordia/Primordia.html)
