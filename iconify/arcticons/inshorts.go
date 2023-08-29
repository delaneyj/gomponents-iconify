package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Inshorts(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="38.5" cy="38.5" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="9.5" cy="38.5" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="38.5" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="38.5" cy="9.5" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.5a4 4 0 0 1 0 8H9.5a4 4 0 0 1 0-8Z"/><circle cx="38.5" cy="23.278" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="23.278" r="4" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}