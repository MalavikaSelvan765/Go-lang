import os
# import responses
import urllib.request as urllib2
import urllib.parse
account_id = os.environ["account_id"]
cookie = os.environ["cookie"]
access_key = os.environ["X_Access_Key_Id"]
access_key_secret = os.environ["X_Access_Key_Secret"]


def get_data():
    url = (f"https://orangescape.kissflow.com/process-report/2/{account_id}/Engineering_Production_Change_Process/PCR_All_Data?apply_preference=true&/page_number=1")
    payload = {}
    headers = {
                "accept": "application/json",
                "X-Access-Key-Id": os.environ["access_key_id"],
                "X-Access-Key-Secret": os.environ["access_key_secret"],
                "Content-Type": "application/json",
                "Cookie": os.environ["cookie"],
            }

    resp = urllib2.Request(url, headers=headers, data=payload)
    response = urllib.request.urlopen(resp)
    return response


print(get_data())

