import time
from functools import wraps


def timer(func):
    @wraps(func)
    def wrapper(*args, **kwargs):
        s = time.time()
        res = func(*args, **kwargs)
        e = time.time()
        print(e - s)
        return res

    return wrapper
