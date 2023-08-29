package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersianCalendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="32.5" cy="11" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="15.5" cy="11" r="2.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.5 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Zm6 29.799V20.701"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.45 20.7v5.252a3 3 0 0 1-3 3H13.5M29.693 20.7H34.5v10.283h0h-4.807a3 3 0 0 1-3-3v-4.282a3 3 0 0 1 3-3ZM34.5 35.299v-4.316"/>`),
		g.Group(children),
	)
}