import unittest

class SegmentDisplay:
    def show(self, value):
        return [["._."],
                ["|.|"],
                ["|_|"]]

class SingleDigitTests(unittest.TestCase):
    def test_zero(self):
        display = SegmentDisplay()
        self.assertEqual(display.show(0), [["._."],
                                           ["|.|"],
                                           ["|_|"]])

unittest.main()
