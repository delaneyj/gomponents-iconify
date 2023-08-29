package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M618.75 99.202v178.006H0V99.202h618.75zm328.125 274.53v178.006H0V373.732h946.875zM731.25 648.262v178.006H0V648.262h731.25zM1200 922.792v178.006H0V922.792h1200z"/>`),
		g.Group(children),
	)
}