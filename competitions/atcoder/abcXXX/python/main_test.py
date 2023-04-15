from main import is_prime, quick_sort, get_primes


def test_is_prime():
    """
    素数判定をテスト
    """
    assert not is_prime(1)
    assert is_prime(2)
    assert is_prime(3)
    assert not is_prime(4)
    assert is_prime(5)
    assert not is_prime(6)
    assert is_prime(7)
    assert not is_prime(8)
    assert not is_prime(9)
    assert not is_prime(10)


def test_quick_sort():
    """
    Tests the quick_sort function to ensure it correctly sorts a list of integers in ascending order.
    """
    arr = [4, 2, 5, 1, 3]
    assert quick_sort(arr) == [1, 2, 3, 4, 5]

    arr = []
    assert quick_sort(arr) == []

    arr = [1]
    assert quick_sort(arr) == [1]

    arr = [5, 4, 3, 2, 1]
    assert quick_sort(arr) == [1, 2, 3, 4, 5]

    arr = [1, 3, 2, 5, 4]
    assert quick_sort(arr) == [1, 2, 3, 4, 5]

    arr = [7, 2, 9, 1, 8]
    assert quick_sort(arr) == [1, 2, 7, 8, 9]

def test_get_primes():
    # 一般的なケース
    assert get_primes(30) == [2, 3, 5, 7, 11, 13, 17, 19, 23, 29]
    
    # Nが素数の場合
    assert get_primes(23) == [2, 3, 5, 7, 11, 13, 17, 19, 23]

    # Nが1の場合
    assert get_primes(1) == []

    # Nが0の場合
    assert get_primes(0) == []

    # Nが負数の場合
    assert get_primes(-5) == []
