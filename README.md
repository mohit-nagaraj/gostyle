# Fortune Cow Rainbow

Fortune Cow Rainbow is a fun project that generates a random fortune, formats it with a cow image, colors it with RGB values, and prints it to the console.
![image](https://github.com/user-attachments/assets/2325b093-5f9f-4464-955c-023267752129)

### Project Structure
* main.go: The main program logic to generate a random fortune.
* fortunes/: A directory containing fortune files.
* cowsay/cow.go: A package to create text balloons and display a cow image.
* catlol/cat.go: A package to color the text with a rainbow effect.

## Usage
### Running the Project
You can run the project with the following command:

```
go run main.go
```
<br>

### Building and Installing
To build and install the project, follow these steps:

1. Clone the repository:

    ```
    git clone https://github.com/mohit-nagaraj/gowow.git
    cd fortune-cow-rainbow
    ```
2. Build the project:

    ```
    go build -o gowow main.go
    ```
Run the built executable

```main.go```

The main.go file reads fortune files from the fortunes directory, selects a random fortune, and prints it. If the specified filename is a directory, it reads all files from the directory, checks if they are fortune files, and selects one randomly to print.

```cowsay/cow.go```

This file creates a text balloon with the fortune and appends a cow image below the text. It processes the input text, adds borders, and normalizes the string lengths to ensure proper alignment within the balloon.

```catlol/cat.go```

This file colors the input text using a rainbow effect. It calculates the RGB values for each character and prints the colored text to the console.

## Contributing
Feel free to fork the repository, create a branch, and submit a pull request with your improvements.

## License
This project is licensed under the Apche License - see the LICENSE file for details.
