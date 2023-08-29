package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoAngular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M13.5 16h4.9L16 10.8L13.5 16z"/><path fill="currentColor" d="M16 3L3 7.6l2.7 15.8L16 29l10.3-5.6L29 7.6L16 3zm5.1 18.6l-1.5-3.2h-7.1L11 21.6H8.6L16 5.3l7.4 16.2h-2.3z"/>`),
		g.Group(children),
	)
}