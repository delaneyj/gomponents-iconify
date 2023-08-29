package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reolink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.364 24.079v7.4h7.4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.364 24.079a7.4 7.4 0 0 1 7.4-7.4h0a7.4 7.4 0 0 1 7.4 7.4h0a7.4 7.4 0 0 1-7.4 7.4h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 24A18.5 18.5 0 0 1 24 5.5h0A18.5 18.5 0 0 1 42.5 24h0A18.5 18.5 0 0 1 24 42.5h0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 24v18.5H24"/>`),
		g.Group(children),
	)
}