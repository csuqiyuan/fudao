import requests
import redis
import json
from hyper.contrib import HTTP20Adapter


def get_subject_headers(grade, subject):
    headers = {
        ":authority": "fudao.qq.com",
        ":method": "GET",
        ":path": "/cgi-proxy/course/discover_subject?client=4&platform=3&version=30&grade={}&subject={}&showid=0&page=1&size=10".format(
            grade, subject),
        ":scheme": "https",
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36",
        "Referer": "https://fudao.qq.com/"
    }
    return headers


def get_grade():
    headers = {
        ":authority": "fudao.qq.com",
        ":method": "GET",
        ":path": "/cgi-proxy/course/grade_subject",
        ":scheme": "https",
        "User-Agent": "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36",
        "Referer": "https://fudao.qq.com/"
    }
    return headers


def get_data(url, headers):
    sessions = requests.session()
    sessions.mount('https://fudao.qq.com', HTTP20Adapter())
    r = sessions.get(url, headers=headers)
    data = str(r.content, 'utf8')
    return data


def redis_set(key, value):
    conn = redis.Redis(host="localhost", port=6379, password="")
    conn.set(key, value)


def save_grades(url):
    grades = get_data(url, get_grade())
    print(grades)
    redis_set("grades", grades)


def save_subjects(url, grades):
    grades = json.loads(grades)
    grades = grades["result"]["grade_subjects"]
    count = 0
    for grade in grades:
        subjects = grade["subject"]
        for subject in subjects:
            data = get_data(url, get_subject_headers(grade["grade"], subject))
            count = count + 1
            print(data)
            key = "{}_{}".format(grade["grade"], subject)
            redis_set(key, data)
    return count

#
url1 = "https://fudao.qq.com/cgi-proxy/course/discover_subject"
url2 = "https://fudao.qq.com/cgi-proxy/course/grade_subject"
#
print(save_subjects(url1, get_data(url2, get_grade())))
save_grades(url2)
