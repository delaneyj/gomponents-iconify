package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Proxdroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 8.796h0a8.262 8.262 0 0 0-11.673.522L24 17.878l-7.827-8.56A8.262 8.262 0 0 0 4.5 8.796h0L18.402 24L4.5 39.205h0a8.262 8.262 0 0 0 11.673-.523L24 30.122l7.827 8.56a8.262 8.262 0 0 0 11.673.523h0L29.598 24Z"/>`),
		g.Group(children),
	)
}