package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func E(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.138.053a12.018 12.018 0 0 0-9.646 3.481A12.014 12.014 0 0 0 .937 16.651a12.014 12.014 0 0 0 11.162 7.348a2.275 2.275 0 1 0-.037-4.549a7.438 7.438 0 0 1-6.932-4.562a7.438 7.438 0 0 1 1.587-8.145a7.437 7.437 0 0 1 12.378 3.014H12.05a2.275 2.275 0 1 0 0 4.55h9.674A2.275 2.275 0 0 0 24 12.04A12.014 12.014 0 0 0 16.597.914a11.962 11.962 0 0 0-3.459-.86Z"/>`),
		g.Group(children),
	)
}