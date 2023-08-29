package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrivatTerminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 28.625h13.875V42.5H5.5z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 23.075V5.5h37v37H24.925c0-14.8-4.625-19.425-19.425-19.425Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.9 16.6v-3.7h22.2v22.2h-3.7c0-11.1-9.25-18.5-18.5-18.5Z"/>`),
		g.Group(children),
	)
}