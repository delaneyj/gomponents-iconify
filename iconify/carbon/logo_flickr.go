package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoFlickr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="8" cy="16" r="6" fill="currentColor"/><circle cx="24" cy="16" r="6" fill="currentColor"/>`),
		g.Group(children),
	)
}