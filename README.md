# Gmail Reader

## Overview

Gmail Reader is a Go script designed to mark emails as read. Nothing serious, just something I hacked together a couple of months ago to mark all my unread emails as read. I thought it might be useful to someone else, so I decided to share it.

## Prerequisites

Before using the script, ensure you have set up the necessary credentials:

1. Go to the [Google Cloud Console](https://console.cloud.google.com/).
2. Create or select a project.
3. Enable the Gmail API.
4. Create OAuth 2.0 Client ID credentials.
5. Download the credentials as `credentials.json` and place it in the project directory.

## Usage

1. **Download the Binary:**

   - [Windows (64-bit)](bin/windows/gmail-reader.exe)
   - [macOS (64-bit)](bin/mac/gmail-reader)
   - [Linux (64-bit)](bin/linux/gmail-reader)

2. **Authentication:**

   - For the first run, follow on-screen instructions to authenticate and mark emails as read.
   - To switch accounts, delete `token.json` and re-run the script.

### Extracting the Authentication Code

1. **Run the Script:**
   - On Windows: `.\gmail-reader.exe`
   - On macOS/Linux: `./gmail-reader`

2. **Authorization Page:**
   - Opened in your default browser, log in to the Gmail account to grant permissions.

3. **Grant Permissions:**
   - Click "Allow" to grant necessary permissions.

4. **Extract the Code:**
   - Copy the authentication code from the URL after granting permissions.

5. **Paste the Code:**
   - Return to the terminal and paste the code when prompted.

6. **Complete Authentication:**
   - Press Enter to finish the process.

7. **Continue Script Execution:**
   - The script proceeds to mark emails as read.

## Troubleshooting: Invalid Grant Error

If encountering an "invalid grant" error:

1. **Check Credentials:**
   - Verify `credentials.json` correctness.

2. **Generate New Credentials:**
   - In [Google Cloud Console](https://console.cloud.google.com/), recreate OAuth credentials.

3. **Verify OAuth 2.0 Consent Screen:**
   - Ensure proper configuration with required scopes.

4. **Clear Cached Tokens:**
   - Delete `token.json` and re-run the script.

5. **Check System Date and Time:**
   - Ensure accuracy; incorrect date/time can affect OAuth token validation.

6. **Verify Internet Connection:**
   - Stable internet is required for communication with Google's servers.

7. **Regenerate OAuth Token:**
   - Follow on-screen instructions during authentication.

If issues persist, open an issue in this repository.

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.