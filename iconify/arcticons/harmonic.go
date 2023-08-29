package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Harmonic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 24.452s18.03-28.31 39-.432"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 24.02c-18.292 43.075-20.713-43.01-39 .432c10.75-34.832 12.645 12.14 19.198 12.183c8.159.054 9.35-47.193 19.802-12.615Z"/>`),
		g.Group(children),
	)
}