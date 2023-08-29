package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LighthousePm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="24" height="39" x="12" y="4.5" fill="none" stroke="currentColor" rx="2" ry="2"/><circle cx="24" cy="10" r="2" fill="none" stroke="currentColor"/>`),
		g.Group(children),
	)
}