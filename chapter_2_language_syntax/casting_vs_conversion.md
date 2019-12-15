# TYPE CASTING
With casting, the compiler pretends that a particular type is really someother type.
It overlays the desired type over the actual type and allows/restricts access to certain bytes. 
No new values are created.
There is a risk of data corruption as we're meddling with bytes that might not have been intended to be meddled with. 

# CONVERSION
With conversion, the compiler does no overlaying of anykind. 
It converts a whole new sequence of bytes of the desired type and works with those bytes instead.
New values are created.
Original value is untouched.

GO advocates CONVERSION over CASTING. It does not ship with default casting behavior - it can be used via a package. 
