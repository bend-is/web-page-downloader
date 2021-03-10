# web-page-downloader

* Downloads the specified web page
* Prints statistics of unique words count

## Usage

Clone the repository:

    $ git clone https://github.com/bend-is/web-page-downloader.git
    $ cd web-page-downloader

Build the application:

    $ go build -o web-page-downloader cmd/web-page-downloader/main.go 

Run the application by specifying the url for a download using the `-u` flag:

    $ ./web-page-downloader -u https://github.com/

Run `./web-page-downloader --help` to see other options