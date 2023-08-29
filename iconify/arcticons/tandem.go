package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tandem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.313 11.813H43.5V24a12.188 12.188 0 0 1-12.188 12.188h0A12.188 12.188 0 0 1 19.126 24v0a12.188 12.188 0 0 1 12.188-12.188Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.688 36.188H4.501h0V24a12.188 12.188 0 0 1 12.188-12.188h0A12.188 12.188 0 0 1 28.876 24h0a12.188 12.188 0 0 1-12.188 12.188Z"/>`),
		g.Group(children),
	)
}