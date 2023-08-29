package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarGoogleThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.731 24a2.257 2.257 0 0 1 2.25 2.25h0a2.257 2.257 0 0 1-2.25 2.25h-.9a3.66 3.66 0 0 1-2.812-.787m0-7.425a3.828 3.828 0 0 1 2.812-.788h.9a2.257 2.257 0 0 1 2.25 2.25h0a2.257 2.257 0 0 1-2.25 2.25h-2.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h26l9-9v-26a2 2 0 0 0-2-2Zm-7 0v37m-19-28v28m28-9h-37m37-19h-28"/>`),
		g.Group(children),
	)
}