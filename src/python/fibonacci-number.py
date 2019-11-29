class Solution:
    def fib(self, N: int) -> int:
        result_list = [1,1]
        if N == 0:
            return 0
        elif N < 2:
            return 1
        else:
            for i in range(2, N):
                result_list[0], result_list[1] = result_list[1], result_list[0] + result_list[1]
        return result_list[-1]

if __name__ == '__main__':
    fibo = Solution()
    print(fibo.fib(40))
