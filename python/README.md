# Python

## Запуск

```
# установка зависимостей
pip install --no-cache-dir -r requirements.txt

FLASK_APP=skill.py flask run
```

## Листинг

```python
# coding: utf-8
from __future__ import unicode_literals
import json

from flask import Flask, request
app = Flask(__name__)

@app.route('/', methods=['POST'])
def handler():
  response = {
    "response": {},
    "session": request.json['session'],
    "version": request.json['version']
  }

  req_text = request.json['request']['original_utterance']
  response['response']['text'] = 'hello' if len(req_text) == 0 else req_text

  return json.dumps(
    response,
    ensure_ascii=False,
    indent=2
  )
```
