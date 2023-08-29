package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HelpSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-miterlimit="10" stroke-width="40" d="M160 164c0-10 1.44-33 33.54-59.46C212.6 88.83 235.49 84.28 256 84c18.73-.23 35.47 2.94 45.48 7.82C318.59 100.2 352 120.6 352 164c0 45.67-29.18 66.37-62.35 89.18S248 290.36 248 316"/><rect width="56" height="56" x="220" y="368" fill="currentColor" rx="3.5" ry="3.5"/>`),
		g.Group(children),
	)
}