package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HintonPlot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M2 2h4v4H2z" fill="currentColor"/><path d="M10 2h4v4h-4z" fill="currentColor"/><path d="M18 2h4v4h-4z" fill="currentColor"/><path d="M26 2h4v4h-4z" fill="currentColor"/><path d="M2 10h4v4H2z" fill="currentColor"/><path d="M10 10h4v4h-4z" fill="currentColor"/><path d="M18 10h4v4h-4z" fill="currentColor"/><path d="M26 10h4v4h-4z" fill="currentColor"/><path d="M2 18h4v4H2z" fill="currentColor"/><path d="M10 18h4v4h-4z" fill="currentColor"/><path d="M18 18h4v4h-4z" fill="currentColor"/><path d="M26 18h4v4h-4z" fill="currentColor"/><path d="M2 26h4v4H2z" fill="currentColor"/><path d="M10 26h4v4h-4z" fill="currentColor"/><path d="M18 26h4v4h-4z" fill="currentColor"/><path d="M26 26h4v4h-4z" fill="currentColor"/>`),
		g.Group(children),
	)
}