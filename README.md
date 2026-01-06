# TCC - Estudando autômatos celulares com a linguagem de programação Go

## TCC - Studying cellular automata using the Go programming language

_This README is still a work in progress_

_This README includes acknowledgements, details on the repo, details on each simulation, and instructions._

### ACKNOWLEDGEMENTS
First and foremost, special thanks are due to Professor Dr. Breno Ferraz de Oliveira for his overall guidance throughout the development of this project, including valuable insights, suggestions, corrections, and code fixes.

Special thanks to Luiz Felipe Locatelli Giroldo for a myriad of reasons.

This section acknowledges the various online sources, like blogs, websites, repositories, and previous works, that informed the development of the three cellular automata simulations (Game of Life, SmoothLife, and Lenia). These materials supported the project by contributing to the research, deepening understanding, or offering inspiration and insight. We sincerely thank the creators and contributors of these resources.

All literature to first understand the three cellular automata and a few derivatives, including their original papers, is referenced in the TCC. We acknowledge and thank these as well.

#### GENERAL
- My and Professor Oliveira's work on C and Go codes for a May-Leonard RPS (rock-paper-scissors) model [(work is "Relatório_Final_Go_C.pdf", download to read)](Relatório_Final_Go_C.pdf) [(Go code available under /Code)](Code/may_leonard_rps.go)

#### SMOOTHLIFE
- [duckythescientist's SmoothLife Python code](https://github.com/duckythescientist/SmoothLife/)
- [ionreq's SmoothLife application](https://sourceforge.net/projects/smoothlife/)
- [ionreq's YouTube channel](https://www.youtube.com/channel/UC_xsxCHaz_h-GGtOaFRGjvg)
- [mikolalysenko (from 0fps.net)'s text on SmoothLife](https://0fps.net/tag/smoothlife/)
- [Rudyon's SmoothLife raylib-go code](https://github.com/rudyon/smoothlife-go/)
- [tsoding's SmoothLife C code](https://github.com/tsoding/SmoothLife/tree/master)

#### LENIA
- [Bert Wang-chak Chan (Chakazul)'s Lenia portal](https://chakazul.github.io/lenia.html)
- [Bert Wang-chak Chan's Lenia coding tutorial](https://colab.research.google.com/github/OpenLenia/Lenia-Tutorial/blob/main/Tutorial_From_Conway_to_Lenia.ipynb)
- [Bert Wang-chak Chan (Chakazul)'s Lenia code repository](https://github.com/Chakazul/Lenia)
- [Bert Wang-chak Chan (Chakazul)'s online Lenia simulator, with SmoothLife and GoL variations](https://chakazul.github.io/Lenia/JavaScript/Lenia.html)
- [Bert Wang-chak Chan (Chakazul)'s online Primordia simulator](https://chakazul.github.io/Primordia/Primordia.html)

### THE REPOSITORY
This repository's primary purpose was to house the videos of our simulations and make it more accessible for the reader to visualize the mechanisms of Game of Life, SmoothLife, and Lenia in time. Now it also houses the code files, the plot files, and instructions. The repository includes:
- Updated (early 2025) code files for the three cellular automata simulations (Game of Life, SmoothLife, and Lenia) in Go language.
- Results in video (.mp4) using specific initial conditions;
- Plot files (in "plt" folder) to generate .png files of the whole grid for every generation and .pdf files of population density curves;
- TCC (.pdf file. PT-BR only).

### DETAILS

_This section does not discuss GoL, SL, and Lenia in depth as cellular automata; rather, we discuss the coding processes and general and particular mechanisms._

#### GENERAL
- The overall structure for the three cellular automata was borrowed from our Go code for the May-Leonard RPS (rock-paper-scissors) model. Because it is essentially a lattice evolving in time, the most primitive "skeleton" for the cellular automata was already present.
- No external libraries were used in any of the three simulations. There is no need for any additional downloads or setup.
- All three operate under the functions ic (initial condition, "generation zero") and op (to create, open, write, and close data files regarding the states of each cell and the population density for every generation). Their basic structures are also similar: ic is run, the ic data file is created, neighborhoods are calculated, the automaton's core equations are used to orderly evolve each cell, the grid is updated, an updated data file is created, and the simulation moves on to the next generation. The simulations use boundary conditions.

#### GAME OF LIFE
- The simulation for Conway's Game of Life was written as an adaptation of our RPS code, altering initial conditions, neighborhood calculations, and overall cell behaviors.
- The code basically uses ic and op and the eight neighbors' equations to determine and register whether cells are dead (0) or alive (1) in a generation.
- Many tested parameters were found in websites dedicated to GoL's patterns.

#### SMOOTHLIFE
- The SmoothLife simulation was first an adaptation of our Game of Life simulation, then rewritten following Rafler's paper.
- The differences to Conway's Game of Life lie in the neighborhood mechanisms (simple Moore neighborhood vs. inner and outer neighborhoods with fillings), the equations used to define a cell's next state, and the grid update mechanism.
- A few tested parameters were extracted directly from ionreq's SmoothLife application, from a list of parameter values.

#### LENIA
- For Lenia, the code was written using Chan's extensive work as a guide, save the obligatory ic and op functions. We'd like to thank Chan for making many helpful Lenia resources easily accessible, such as pseudocode, Python code, online Javascript simulation and coding tutorials.
- There was an attempt on earlier versions of the simulation to implement a Fast Fourier Transform (FFT) algorithm, using Go FFT libraries such as [argusdusty's](https://github.com/argusdusty/gofft) and [mjibson's](https://pkg.go.dev/github.com/mjibson/go-dsp/fft). However, FFT implementation was scrapped.  
- A 2D manual convolution method is written as a function, based on well-known convolution methods, as it is a required step according to Chan's paper.
- ##### Testing of this simulation was rather limited. The results obtained thus far indicate that Lenia species can form in this simulation, but tend to disappear in a few generations. If fixes are necessary, the code will be swiftly updated.

### INSTRUCTIONS
- You can run these simulations if Go is installed on your computer.
- Download the simulations (.go files). You may need to move the files after download to go/src on your computer. Otherwise, it's recommended to create a separate folder and move the .go files so that later data files are neatly stored;
- Open the terminal (Windows cmd, Cygwin, etc.), move to the directory, and type in "go run simulation-name.go" (adjust "simulation-name" accordingly);
- Each simulation will produce a number of .dat files (specified as NG, number of generations). These .dat files contain the entire grid, each number imprinted representing the current state of the cell at the specified position;
- To generate grid images from the .dat files, Gnuplot is necessary.
- Download and edit the grid.plt file (number of generations, grid size, .dat files names, output files names, etc.). Ensure it is in the same directory as the .dat files and run it;
      - The grid.plt file contains a rainbow color palette to plot SmoothLife and Lenia simulations. To plot Conway's Game of Life, you may want to adjust it to only two colors.
- The same printing process can be done using the pop-density.plt file and the population density .dat file generated by the simulation. In this case, one .pdf graph can display many data files (curves) at once;
- To make videos from the .png images, open the terminal after generating all images and ensure you are in the same directory. Then, copy and paste the following command, replacing "image-%d.png" and "video.mp4" as necessary:

ffmpeg -r 7 -i image-%d.png -c:v libx264 -profile:v baseline -level 3.0 -pix_fmt yuv420p -y video.mp4

- Parameters may be edited in the .go files. They are generally presented as constants, variables, equations or functions.
- There are two initial conditions to be used in these simulations. One is a pseudorandom state distribution through the entire grid (the standard beginning for GoL), and the other is the random spawning of squares containing pseudorandom decimal numbers. The latter initial condition is a better option for SmoothLife and Lenia.
