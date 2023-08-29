package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignJustify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 99.202v178.006h1200V99.202H0zm0 274.53v178.006h1200V373.732H0zm0 274.53v178.006h1200V648.262H0zm0 274.53v178.006h1200V922.792H0z"/>`),
		g.Group(children),
	)
}