package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outbank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" d="M13.5 42.5c3 0 17-4 17-13m-17-24c3 0 17 4 17 13v11"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M34.5 42.5c-3 0-17-4-17-13m17-24c-3 0-17 4-17 13v11"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}