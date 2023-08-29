package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Maxcdn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28.819 27.669h-6.088L25.975 12.5c.144-.637.056-1.219-.275-1.606c-.313-.381-.856-.6-1.512-.6h-3.081l-3.719 17.375h-6.087l3.719-17.375H9.807L6.088 27.669H.001L3.72 10.294L.932 4.332h23.256c2.462 0 4.706 1.019 6.144 2.806c1.456 1.788 1.988 4.213 1.475 6.619z"/>`),
		g.Group(children),
	)
}