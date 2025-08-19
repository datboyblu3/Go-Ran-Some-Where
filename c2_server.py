import logging
from http.server import SimpleHTTPRequestHandler, HTTPServer
import threading
import urllib.parse

#Disable all logs from http.server
class NoLogRequestHandler(SimpleHTTPRequestHandler):
    def log_message(self, format, *args):
        pass

#Configure logging to print only errors
logging.basicConfig(level=logging.ERROR)
                                   
class MyRequestHandler(NoLogRequestHandler):
    def do_GET(self):
        global user_input
        # Parse the data parameter from the URI
        query_params = urllib.parse.parse_qs(urllib.parse.urlparse(self.path).query)
        data_param   = query_params.get('data', [''])[0]

        #Print the entered command and the received "data" parameter
        if data_param != "":
            print(f"\nData parameter received: {data_param}")

        #Send the user input as response
        self.send_response(200)
        self.send_header('Content-type', 'text/plain')
        self.end_headers()
        self.wfile.write(bytes(user_input, 'utf-8'))

        user_input = ''


def get_user_input():
    global user_input
    while True:
        user_input = input("Command (type 'exit' to stop) > ")
        if user_input.lower() == 'exit':
            break


if __name__ == '__main__':
    try:
        port = 9001
        server_address = ('', port)
        user_input = ''

        #Create an HTTP server with custom request handler
        httpd = HTTPServer(server_address, MyRequestHandler)

        #Start server in a separate thread
        server_thread = threading.Thread(target=httpd.server_forever)
        server_thread.start()

        #Start user input
        get_user_input()

        #Stop server when user input loop is exited
        httpd.shutdown()
        httpd.server_close()

        print("Server stopped")

    except KeyboardInterrupt:
        httpd.shutdown()
        httpd.server_close()
        print("\nServer stopped")