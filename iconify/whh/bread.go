package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bread(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M769.07 629q18 17 67 8t90-39q50 79 74.5 148.5t23.502 119.5t-13.5 84.5t-35.5 56.5q-46 47-145.5-30.5T500.07 668l-72-72q-241-241-289-292q-137-149-139-212q-1-25 17-43q21-21 53-33.5t77.5-15t109 17T392.07 78q-8 11-11 14q-26 30-60.5 54t-63.5 35q11 9 33.5 10.5t48.5-5t53.5-22t47.5-37.5q0-1 9-13q93 62 197 159q-1 0-4 4.5t-5 6.5q-27 30-61 54t-63 35q11 9 33.5 10.5t48-5t53-22t47.5-36.5l2.5 2.5l2.5 2.5q113 113 189 219q-59 63-120 85z"/>`),
		g.Group(children),
	)
}