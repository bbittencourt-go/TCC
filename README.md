#TCC | Code from my TCC (Game of Life, SmoothLife, Lenia) in Go language

Details:
- No additional libraries used
- Manual convolution (where applicable, written as a function)

Instructions:
- If you already have Go installed on your computer, you can run these simulations. If not, install Go at https://go.dev/.
- Download the simulations (.go files). You may need to move the files after download to go/src.
- Open a terminal (Windows cmd, Cygwin, etc.), move to the directory where the simulations are and type in "go run (simulationname).go"
- Each simulation will produce a finite amount of .dat files (specified in the code as the constant NG). These .dat files contain the entire grid, with each number representing the current state of the cell at the specified position.
- To generate images of the grid, you will need Gnuplot. Make sure you have it in your computer or download it at http://www.gnuplot.info/download.html.
- Download and edit the grid.plt file to your specifications (number of generations, grid size, file name etc.). Make sure it is in the same directory as the .dat files and execute it by clicking on it.
- To make a video, open the terminal after generating all images, making sure to select the same directory as them. Then, copy and paste the following command, replacing "image-%d.png" and "video.mp4" as necessary: 
ffmpeg -r 7 -i image-%d.png -c:v libx264 -profile:v baseline -level 3.0 -pix_fmt yuv420p -y video.mp4
