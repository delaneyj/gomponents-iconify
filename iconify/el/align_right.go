package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M581.25 99.202v178.006H1200V99.202H581.25zm-328.125 274.53v178.006H1200V373.732H253.125zm215.625 274.53v178.006H1200V648.262H468.75zM0 922.792v178.006h1200V922.792H0z"/>`),
		g.Group(children),
	)
}