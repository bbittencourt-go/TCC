# TCC - Estudando autômatos celulares com a linguagem de programação Go

## TCC - Studying cellular automata using the Go programming language

_this README is still work in progress_
_this README includes acknowledgements, details on the repo, details on each simulation and instructions._

### ACKNOWLEDGEMENTS
First and foremost, special thanks to Professor Dr. Breno Ferraz de Oliveira for his overall guidance during the development of this project, including valuable insights, suggestions, corrections, and small code fixes.

Special thanks to Luiz Felipe Locatelli Giroldo for a myriad of reasons.

This section acknowledges the various sources, like articles, blogs, websites, repositories, and previous works, that informed the development of the three cellular automata simulations (Game of Life, SmoothLife, and Lenia). These materials supported the project by contributing to the research, deepening understanding, or offering inspiration and insight. We sincerely thank the creators and contributors of these resources.

- My and Professor Oliveira's work on C and Go codes for a May-Leonard RPS (rock-paper-scissors) model [(work is "Relatório_Final_Go_C.pdf", download to read)](Relatório_Final_Go_C.pdf) [(Go code available under /Code)](Code/may_leonard_rps.go) 
- [Bert Wang-chak Chan (Chakazul)'s Lenia portal](https://chakazul.github.io/lenia.html)
- [Bert Wang-chak Chan's Lenia coding tutorial](https://colab.research.google.com/github/OpenLenia/Lenia-Tutorial/blob/main/Tutorial_From_Conway_to_Lenia.ipynb)
- [Bert Wang-chak Chan (Chakazul)'s Lenia code repository](https://github.com/Chakazul/Lenia)
- [Bert Wang-chak Chan (Chakazul)'s online Lenia simulator, with SmoothLife and GoL variations](https://chakazul.github.io/Lenia/JavaScript/Lenia.html)
- [Bert Wang-chak Chan (Chakazul)'s online Primordia simulator](https://chakazul.github.io/Primordia/Primordia.html)
- [ionreq (Stephan Rafler)'s SmoothLife application](https://sourceforge.net/projects/smoothlife/)
- [ionreq (Stephan Rafler)'s YouTube channel](https://www.youtube.com/channel/UC_xsxCHaz_h-GGtOaFRGjvg)
- [Rudyon's SmoothLife raylib-go code](https://github.com/rudyon/smoothlife-go/tree/9fa85b9a457a2d817529fdfa9b8e062352d516fd)
- [mikolalysenko (from 0fps.net)'s text on SmoothLife](https://0fps.net/tag/smoothlife/)

### THE REPOSITORY
This repository's primary purpose was to house the videos of my simulations and make it more accessible for the reader to visualize the mechanisms of Game of Life, SmoothLife and Lenia in every generation. Now it also houses the code files, the plot files and these instructions for whoever is interested. The repository includes:
- Updated code files for the three cellular automata simulations (Game of Life, SmoothLife and Lenia) in Go language.
- Results in video (.mp4) using specific initial conditions (download to watch!);
- Plot files (in "plt" folder) to generate .png files of the whole grid for every generation and .pdf files of population density curves;
- TCC (.pdf file. PT-BR only).

### DETAILS

#### GENERAL
- The overall structure of the code for the three cellular automata was heavily borrowed from my first code project, also assisted by Professor Oliveira, which was Go code for the May-Leonard RPS (rock-paper-scissors) model (http://www.eaic.uem.br/eaic2023/anais/artigos/6438.pdf).
- No external libraries are used in any of the three simulations, meaning there is no need for any downloads.
- Functions ic (initial condition, "generation zero") and op (to create files regarding the states of each cell and the population density for every generation).

#### GAME OF LIFE
For Conway's Game of Life, most of the structure was written somewhat as an adaptation of my original RPS code, altering the initial condition, the neighborhood calculations and the overall behavior for every cell.

#### SMOOTHLIFE
- The SmoothLife simulation was primarily written as a sort of evolution of the Game of Life simulation, which is also true coding aside.
- The crucial differences lie in the neighborhood mechanisms (simple Moore neighborhood vs. inner and outer neighborhoods with fillings), the equations used to define a cell's next state (all presented by Rafler in his paper) and the grid update mechanism.
- Whilst SmoothLife was not the most difficult simulation in this project, it took the longest to write due to many errors, bugs and mistakes, especially when it came to updating the cells' states. 

#### LENIA
- The code was mostly inspired by Bert Wang-chak Chan's work such as the pseudocode in his paper, his Lenia Python code, the online Lenia simulation and his two Lenia coding tutorials. We'd like to specially thank Chan for how accessible these resources are!
- A 2D manual convolution method is written as a function, based on well-known convolution methods, as it is a required step according to Chan's paper.
- There was an attempt on earlier versions of the simulation to implement Fast Fourier Transform (FFT) as a way to accelerate execution times, as suggested by Professor Oliveira. However, implementing it proved to be a difficult task (both writing by hand and using Go FFT libraries available online such as [argusdusty's](https://github.com/argusdusty/gofft) and [mjibson's](https://pkg.go.dev/github.com/mjibson/go-dsp/fft)). Because of the impressive number of errors, endless debugging and the fact the simulation wasn't even functional by itself at that point, we decided to scrap that and focus on actually having Lenia.

### INSTRUCTIONS
- You can run these simulations if Go is installed on your computer. If not, install Go at https://go.dev/;
- Download the simulations (.go files). You may need to move the files after download to go/src on your computer (in my experience, it was unnecessary). Otherwise, it's recommended to create a separate folder and move the .go files to this destination so that later data files are neatly stored, separate from unrelated files;
- Open the terminal (Windows cmd, Cygwin, etc.), move to the directory where the .go files are stored, and type in "go run simulation-name.go" (adjust "simulation-name" accordingly);
- Each simulation will produce a finite number of .dat files (specified in the code as the constant NG, number of generations). These .dat files contain the entire grid, each number imprinted representing the current state of the cell at the specified position;
- To generate grid images from the .dat files, you will need Gnuplot. Ensure it is on your computer or download it at http://www.gnuplot.info/download.html;
- Download and edit the grid.plt file to your specifications (number of generations, grid size, .dat files names, output files names, etc.). Ensure it is in the same directory as the .dat files and run it by simply clicking;
      - The grid.plt file contains a rainbow color palette to plot SmoothLife and Lenia simulations (as these use decimal numbers). To plot Conway's Game of Life, you may want to adjust it to only two colors.
- The same printing process can be done using the pop-density.plt file and the population density .dat file generated by the simulation. In this case, one .pdf graph can display many data files (curves) at once, so adjust the file according to your intentions;
- To make videos from the .png images, open the terminal after generating all images and ensure you are in the same directory. Then, copy and paste the following command, replacing "image-%d.png" and "video.mp4" as necessary:

ffmpeg -r 7 -i image-%d.png -c:v libx264 -profile:v baseline -level 3.0 -pix_fmt yuv420p -y video.mp4

- Parameters may be edited in the .go files. They are generally presented as constants, variables, equations or functions, making them easily alterable.
- There are two initial conditions to be used in these simulations. One is a pseudorandom state distribution through the entire grid (the standard beginning for GoL) and the other is the random spawning of squares containing pseudorandom decimal numbers, whilst the rest of the grid is null. The latter initial condition is a better option for SmoothLife and Lenia, as it does not encourage the growth of something called an "expansive mass" (a sort of amoeba that expands until it takes over the entire grid, effectively hindering interactions and formations of any interesting patterns or behaviors). Therefore, the SmoothLife and Lenia simulations are coded to cater to this initial condition. The user may alter this to their wishes.
