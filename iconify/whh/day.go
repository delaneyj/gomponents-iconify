package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Day(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 374q-97 62-96 137q2 78 96 137q-255 12-137 238q-227-118-238 137q-62-97-137-96q-77 2-137 96q-11-255-238-137Q255 660 0 648q97-62 96-137q-2-77-96-137q255-11 137-237Q364 254 375 0q60 93 137 95q75 2 137-95q11 254 238 137q-118 226 137 237zM512 191q-87 0-161 43T234.5 350.5T192 511t42.5 161T351 788.5T512 831t161-42.5T789.5 672T832 511t-42.5-160.5T673 234t-161-43zm0 576q-106 0-181-75t-75-181t75-181t181-75t181 75t75 181t-75 181t-181 75z"/>`),
		g.Group(children),
	)
}