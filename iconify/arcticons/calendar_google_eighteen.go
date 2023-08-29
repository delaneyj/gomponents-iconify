package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarGoogleEighteen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.795 20.738l2.25-1.238m0 0v9m5.91-4.5h-1.463a2.257 2.257 0 0 0-2.25 2.25h0a2.257 2.257 0 0 0 2.25 2.25h1.463a2.257 2.257 0 0 0 2.25-2.25h0a2.257 2.257 0 0 0-2.25-2.25Zm0 0a2.257 2.257 0 0 0 2.25-2.25h0a2.257 2.257 0 0 0-2.25-2.25h-1.463a2.257 2.257 0 0 0-2.25 2.25h0a2.257 2.257 0 0 0 2.25 2.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h26l9-9v-26a2 2 0 0 0-2-2Zm-7 0v37m-19-28v28m28-9h-37m37-19h-28"/>`),
		g.Group(children),
	)
}