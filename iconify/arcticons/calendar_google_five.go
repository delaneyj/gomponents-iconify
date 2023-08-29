package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarGoogleFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.019 27.712a3.529 3.529 0 0 0 2.7.788h.337a2.978 2.978 0 0 0 2.925-2.925h0a2.978 2.978 0 0 0-2.925-2.925H21.02V19.5h5.962"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h26l9-9v-26a2 2 0 0 0-2-2Zm-7 0v37m-19-28v28m28-9h-37m37-19h-28"/>`),
		g.Group(children),
	)
}