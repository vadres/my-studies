from date import Date
from rectangle import Rectangle

date = Date(12, 12, 2022)
date.format()

rec = Rectangle(7, 6)
rec.area = 7
print(rec.get_area())
