from flask import Flask, request, jsonify

import datetime
import pickle
import os.path
from googleapiclient.discovery import build
from google_auth_oauthlib.flow import InstalledAppFlow
from google.auth.transport.requests import Request

# If modifying these scopes, delete the file token.pickle.
SCOPES = ['https://www.googleapis.com/auth/calendar']

app = Flask(__name__)


def main(summary, start, end,frm,to):
    """Creates a Google Calendar event with the given summary."""
    start_datetime = datetime.datetime.strptime(start, '%Y-%m-%dT%H:%M:%SZ')
    end_datetime = datetime.datetime.strptime(end, '%Y-%m-%dT%H:%M:%SZ')

    start_str = start_datetime.isoformat() + 'Z'
    end_str = end_datetime.isoformat() + 'Z'

    print("-----", start_str)
    print("-----", end_str)

    # Create event dictionary with string formatted dates
    event = {
        'summary': summary,
        'start': {'dateTime': start_str},
        'end': {'dateTime': end_str},
        'attendees':[
            {
                'email':frm,
                
            },
            {
                'email':to
            }

        ],
        "conferenceData": {
            "createRequest": {
                "conferenceSolutionKey": {
                    "type": "hangoutsMeet"
                },
                "requestId": "medcare."
            },
        },
        'reminders': {
        'useDefault': False,
        'overrides': [
          {'method': 'email', 'minutes': 24 * 60},
          {'method': 'popup', 'minutes': 10},
        ],
      }
    }

    creds = None
    if os.path.exists('token.pickle'):
        with open('token.pickle', 'rb') as token:
            creds = pickle.load(token)
    if not creds or not creds.valid:
        if creds and creds.expired and creds.refresh_token:
            creds.refresh(Request())
        else:
            flow = InstalledAppFlow.from_client_secrets_file(
                'client_secret.json', SCOPES)
            creds = flow.run_local_server(port=5001)
        with open('token.pickle', 'wb') as token:
            pickle.dump(creds, token)

    service = build('calendar', 'v3', credentials=creds)

    # Insert the event into the calendar
    event = service.events().insert(calendarId='primary', conferenceDataVersion=1, body=event).execute()

    return event.get('hangoutLink'), event.get('id')


@app.route('/create-event', methods=['POST'])
def create_event():
    """Handler for creating a Google Calendar event."""
    print(f'Request Body : \n{request.json}')
    summary = request.json.get('summary')
    start = request.json.get('start')
    end = request.json.get('end')
    frm = request.json.get('from')
    to = request.json.get('to')

    print("-----",start)
    print("-----", end)
    if summary and start and end:
        hangout_link, event_id = main(summary,start,end,frm,to)
        return jsonify({'hangout_link': hangout_link, 'event_id': event_id}), 200
    else:
        return jsonify({'error': 'Summary is required'}), 400

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001)

