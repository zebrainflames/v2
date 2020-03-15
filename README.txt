## Vec2 ##

This is a simple 2D vector lib for personal use in Go projects.

Supports basic operations like additions, subtractions, normalizing, length calculation etc. for different types.

Note that the receiver methods are mutative to avoid allocations. Non-mutative function options exist too. It's to be
tested whether this offers performance benefits (or any other benefits really) in a real use case as of now.

Use at your own risk, the API is very much in flux for now. Released hopefully properly versioned though.