# we can import modules to use later
import base64, sys

def new_accumulator(initial):
    """This is a function that returns a function (a closure)!"""
    total = [initial]
    def accumulator(x):
        total[0] += x
        print total[0]
    return accumulator


class Greeter:
    """I am a docstring.
    
    This is a class definition with some methods.
    """
    
    def __init__(self, name="Spam"):
        self.name = name
    
    def _greet(self):
        return "Hello {}".format(self.name)
    
    def greet(self):
        print self._greet()
        
    def encoded_greet(self):
        print base64.b64encode(self._greet())


# the following code executes only when our script
# is run directly (not imported by another module)
if __name__ == "__main__":
    g = Greeter()
    g.greet()
    g.encoded_greet()
    print # prints a blank line
    
    acc = new_accumulator(-2)
    acc(0) # prints -2
    acc(1) # prints -1
    acc(2) # prints 1
    acc(5) # prints 6


# this code always execute:
print "--> End of {!r}".format(" ".join(sys.argv))