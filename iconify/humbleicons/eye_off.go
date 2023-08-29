package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 5l16 16M11.148 9.123a3 3 0 0 1 3.722 3.752M8.41 6.878C12.674 4.762 17.267 6.47 21 12c-1.027 1.521-2.119 2.753-3.251 3.696m-2.509 1.59C11.076 19.142 6.631 17.38 3 12c1.01-1.496 2.083-2.713 3.196-3.65"/>`),
		g.Group(children),
	)
}