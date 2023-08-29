package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Geonotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 43.5l-7.626-7.626H8.4V4.5H24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 7.894h-3.2l.905 17.121H24"/><circle cx="24" cy="31.174" r="2.684" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 43.5l7.626-7.626H39.6V4.5H24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 7.894h3.2l-.905 17.121H24"/>`),
		g.Group(children),
	)
}