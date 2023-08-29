package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rustdesk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.132 3.08a21.475 21.475 0 0 0 0 41.84m9.734.001a21.475 21.475 0 0 0 0-41.842"/>`),
		g.Group(children),
	)
}