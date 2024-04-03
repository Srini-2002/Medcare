import sys

from oauth2client import client
from googleapiclient import sample_tools


def main(argv):
    # Authenticate and construct service.
    service, flags = sample_tools.init(
        argv,
        "calendar",
        "v3",
        __doc__,
        __file__,
        scope="https://www.googleapis.com/auth/calendar.readonly",
    )

    try:
        page_token = None
        while True:
            calendar_list = service.calendarList().list(pageToken=page_token).execute()
            for calendar_list_entry in calendar_list["items"]:
                print(calendar_list_entry["summary"])
            page_token = calendar_list.get("nextPageToken")
            if not page_token:
                break

    except client.AccessTokenRefreshError:
        print(
            "The credentials have been revoked or expired, please re-run"
            "the application to re-authorize."
        )

import os

from google.auth.transport.requests import Request
from google.oauth2.credentials import Credentials
from google_auth_oauthlib.flow import InstalledAppFlow
from googleapiclient.discovery import build as __build
from googleapiclient.errors import HttpError

def new_test():
    SCOPES = ["https://www.googleapis.com/auth/spreadsheets","https://www.googleapis.com/auth/gmail.modify"]
    credentials = None
    if os.path.exists("token.json"):
        credentials = Credentials.from_authorized_user_file("token.json",SCOPES)
    if not credentials or not credentials.valid:
        if credentials and credentials.expired and credentials.refresh_token:
            credentials.refresh(Request())
        else:
            flow = InstalledAppFlow.from_client_secrets_file("client_secrets.json",SCOPES)
            credentials = flow.run_local_server(port=5001,redirect_uri_trailing_slash=False)
        with open("token.json","w") as token:
            token.write(credentials.to_json())

if __name__ == "__main__":
    new_test()