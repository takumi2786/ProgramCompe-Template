import time
import sys
import io
from typing import List

# テスト用の入力文字列を記載する
INPUT = """\
5
"""

def debug(func: callable):
    """
    このデコレータをつけるとデバッグモードが有効化されます
    標準入力の差し替えと実行時間の計測を行います
    """
    def wrapper(*args, **kwargs):
        sys.stdin = io.StringIO(INPUT)
        start_time = time.perf_counter()
        res = func(*args, **kwargs)
        end_time = time.perf_counter()
        elapsed_time = end_time - start_time
        print(elapsed_time)
        return res
    return wrapper


@debug
def main():
    """
    メインの処理を記述する
    """
    n = int(input())
    
    if is_prime(n):
        print('Prime')
    else:
        print('Not Prime')


def quick_sort(arr: List[int]):
    """
    Sorts an array of integers in ascending order using the QuickSort algorithm.

    Args:
        arr: A list of integers to be sorted.

    Returns:
        The sorted list of integers.
    
    Example:
        >>> arr = [3, 2, 1, 5, 4]
        >>> quick_sort(arr)
        [1, 2, 3, 4, 5]
    """
    if len(arr) <= 1:
        return arr
    pivot = arr[len(arr)//2]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    return quick_sort(left) + middle + quick_sort(right)

def is_prime(n: int) -> bool:
    if n <= 1:
        return False

    if n == 2:
        return True

    if n % 2 == 0:
        return False

    i = 3

    while i * i <= n:
        if n % i == 0:
            return False

        i += 2
    return True

def get_primes(N :int)->List[int]:
    """
    N以下の素数を求める
    エラトスネテスのふるい
    """
    # 負の数は素数に含まれない
    if N < 0:
        return []
    squere = math.sqrt(N)
    is_prime = [True for i in range(0, N+1)] # 整数iが素数かどうかを格納する
    for i in range(0, N+1):
        if i <= 1:
            is_prime[i] = False
            continue
        # 整数iの倍数に印
        if is_prime[i] == True:
            is_prime[i] = True
            # iよりも大きいiの倍数に印をつける
            # ex: i=2... {4, 6, 8, 10, ...}
            for x in range(2*i, N+1, i): # 2*2, 2*2+2, 6+2, ...
                is_prime[x] = False
        # √(N)以下の整数iに何らかの印がついたら終了
        if squere < i:
            break
    primes = []
    for i in range(2, N+1):
        if is_prime[i]:
            primes.append(i)
    return primes

if __name__ == "__main__":
    main()

# 入力の受け取り
# n = int(input())
# A = list(map(int, input().split(" ")))
