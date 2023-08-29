package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarGoogleTwentyNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.391 22.5a2.966 2.966 0 0 1 2.889-3a2.986 2.986 0 0 1 2.111 5.111c-1.222 1-5 3.889-5 3.889h5.889"/><circle cx="27.609" cy="22.5" r="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.942 27.389a2.729 2.729 0 0 0 2.445 1.111h.222a2.987 2.987 0 0 0 3-3v-3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h26l9-9v-26a2 2 0 0 0-2-2Zm-7 0v37m-19-28v28m28-9h-37m37-19h-28"/>`),
		g.Group(children),
	)
}