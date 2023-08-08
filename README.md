# Website Monitoring Application

This is a simple command-line application written in Go that monitors the availability of websites. It is part of a series of Golang studies that I've been doing in the past months. The application periodically sends HTTP requests to the specified websites and logs their status.

## Usage

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/your-username/website-monitor.git
   ```

2. Navigate to the project directory:

   ```bash
   cd website-monitor
   ```

3. Compile and run the application:

   ```bash
   go run main.go
   ```

4. Follow the on-screen instructions to choose from the available options.

## Features

- Start monitoring websites for availability.
- Display logs of website availability and downtime.
- Exit the program.

## How It Works

1. The application reads a list of website URLs from the `sites.txt` file.
2. It then repeatedly sends HTTP requests to each website to check their status.
3. The application logs the results, recording the date, time, website, and status in the `log.txt` file.

## Notes

- Adjust the `monitoring` constant to specify the number of monitoring cycles.
- The `delay` constant controls the time interval between monitoring cycles.
- The `sites.txt` file should contain one website URL per line.
- Logs are stored in the `log.txt` file.

Feel free to modify and enhance the application according to your needs.

Happy monitoring!
```

