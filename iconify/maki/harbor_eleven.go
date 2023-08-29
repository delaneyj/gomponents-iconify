package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HarborEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.5 0A2.5 2.5 0 0 0 3 2.5a2.493 2.493 0 0 0 1.75 2.371v4.545c-.659-.115-1.338-.375-1.893-.857C2.09 7.89 1.5 6.829 1.5 5A.75.75 0 1 0 0 5c0 2.17.773 3.735 1.873 4.691S4.333 11 5.5 11s2.527-.352 3.627-1.309S11 7.171 11 5c.014-1.014-1.514-1.014-1.5 0c0 1.83-.589 2.89-1.357 3.559c-.555.482-1.234.742-1.893.857V4.875A2.494 2.494 0 0 0 5.5 0zm0 1.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}