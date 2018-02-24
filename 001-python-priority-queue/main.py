from Queue import PriorityQueue
from random import randint

N = 1000


def main():
    q = PriorityQueue(1000)
    for i in range(100):
        q.put((randint(0, 1000), i))

    while not q.empty():
        print(q.get())


if __name__ == '__main__':
    main()
