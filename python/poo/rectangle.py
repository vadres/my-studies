class Rectangle:

    def __init__(self, x, y):
        self.__x = x
        self.__y = y
        self.__area = x * y

    def get_area(self):
        return self.__area
