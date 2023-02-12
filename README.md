# <span style="color:#C0BFEC">**🦔 A program that parses a table and saves it to Google Drive**</span>

## <span style="color:#C0BFEC">***Default Settings:***</span>

* Default output file name: `data.csv`
* Default credentials file name: `keys.json`
* URL to parse: `https://confluence.hflabs.ru/pages/viewpage.action?pageId=1181220999`

You can change the settings in the file: `./config/config.go`

## <span style="color:#C0BFEC">***How to get credentials file:***</span>

1. Go to the [Google API Console](https://console.developers.google.com/).
2. Create a new project.
3. Click **Enable API**. Search for and enable the Google Drive API.
4. Create credentials for a Web Server to access Application Data.
5. Name the service account and grant it a Project Role of Editor.
6. Download the JSON file.
7. Copy the JSON file to your code directory and rename it to `keys.json`.
8. Run the code.
9. Check your Google Drive for the new file.

## <span style="color:#C0BFEC">***How to get a token file:***</span>

1. Run the code.
2. Follow the link in the console.
3. Copy the code from the browser and paste it into the console.


## <span style="color:#C0BFEC">***Enter to run:*** </span>

```
go run cmd/main.go
```