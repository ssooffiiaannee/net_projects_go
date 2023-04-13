#include <iostream>
#include <sys/socket.h>
#include <netinet/in.h>
#include <unistd.h>
#include <string.h>

int main() {
    // Create a TCP socket
    int server_fd = socket(AF_INET, SOCK_STREAM, 0);
    if (server_fd == -1) {
        std::cerr << "Failed to create socket" << std::endl;
        return 1;
    }

    // Set up the server address
    struct sockaddr_in server_addr;
    server_addr.sin_family = AF_INET;
    server_addr.sin_addr.s_addr = INADDR_ANY;
    server_addr.sin_port = htons(8080);

    // Bind the socket to the server address
    if (bind(server_fd, (struct sockaddr*)&server_addr, sizeof(server_addr)) == -1) {
        std::cerr << "Failed to bind socket" << std::endl;
        return 1;
    }

    // Listen for incoming connections
    if (listen(server_fd, 10) == -1) {
        std::cerr << "Failed to listen for connections" << std::endl;
        return 1;
    }
    std::cout << "before loop" << std::endl;
    // Loop forever, handling incoming connections
    while (true) {
        // Accept the next incoming connection
        int client_fd = accept(server_fd, NULL, NULL);
        if (client_fd == -1) {
            std::cerr << "Failed to accept connection" << std::endl;
            continue;
        }

        // Read the request from the client
        char buffer[4096];
        memset(buffer, 0, sizeof(buffer));
        int bytes_received = read(client_fd, buffer, sizeof(buffer));
        if (bytes_received == -1) {
            std::cerr << "Failed to read request" << std::endl;
            close(client_fd);
            continue;
        }

        // Print the request to the console
        std::cout << buffer << std::endl;

        std::string web_page = "<html>\
                                    <head>\
                                        <title> my test page </title>\
                                    </head>\n\
                                    <body>\
                                        <h1> this is good ! </h1>\
                                    </body>";

        // Send a response back to the client
        std::string response = "HTTP/3 200 OK\r\nContent-Length:" + std::to_string(web_page.size()) + "\r\n\r\n" + web_page; //Hello, world!";
        int bytes_sent = write(client_fd, response.c_str(), response.length());
        if (bytes_sent == -1) {
            std::cerr << "Failed to send response" << std::endl;
        }

        // Close the client socket
        close(client_fd);
    }

    // Close the server socket
    close(server_fd);

    return 0;
}

