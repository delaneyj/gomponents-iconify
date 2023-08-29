package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Musicplayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="14.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="21.5" cy="29.25" r="3.5"/><path d="M25 29.25v-14m0 0h5v4h-5z"/></g>`),
		g.Group(children),
	)
}