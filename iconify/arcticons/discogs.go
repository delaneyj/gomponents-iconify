package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Discogs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="24" r="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="currentColor" d="M24.75 24a.75.75 0 1 1-.75-.75a.76.76 0 0 1 .75.75Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 7.5A16.5 16.5 0 0 0 7.5 24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 13a11 11 0 0 0-11 11m11 16.5A16.5 16.5 0 0 0 40.5 24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 35a11 11 0 0 0 11-11"/>`),
		g.Group(children),
	)
}