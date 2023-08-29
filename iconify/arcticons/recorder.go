package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Recorder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5h0a7.61 7.61 0 0 1 7.61 7.61v9.42A7.61 7.61 0 0 1 24 29.14h0a7.61 7.61 0 0 1-7.61-7.61v-9.42A7.61 7.61 0 0 1 24 4.5Zm2.94 28.65v9.71a.64.64 0 0 1-.63.64h-4.62a.64.64 0 0 1-.63-.64v-9.71"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.23 23.79a12 12 0 0 0 23.54 0"/>`),
		g.Group(children),
	)
}