package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archlinux(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M704 640q17 17 63 61t83.5 82t69.5 76q80 141 104 165q-51-51-153.5-105.5T640 843V730q0-64-37.5-109T512 576t-90.5 45T384 730v113q-128 21-230.5 75.5T0 1024q33-33 155.5-257.5T375 329q67 49 137 119q-37-74-122-151Q477 102 502 0h20q28 115 124.5 324.5T855 742q-82-68-151-102z"/>`),
		g.Group(children),
	)
}