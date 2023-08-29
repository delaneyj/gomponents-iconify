package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Topbuzz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4" d="m18.905 18.168l-8.58-1.824a1.032 1.032 0 0 1-.794-1.224l1.931-9.087a1.032 1.032 0 0 1 1.225-.795l29.28 6.223c.558.119.914.667.796 1.225l-1.932 9.087a1.032 1.032 0 0 1-1.224.795l-8.506-1.808c-.388-.082-.617.046-.69.385l-3.97 18.676a1.032 1.032 0 0 1-1.223.795l-9.088-1.931a1.032 1.032 0 0 1-.795-1.225l3.971-18.681c.071-.335-.062-.54-.4-.611Z"/>`),
		g.Group(children),
	)
}