package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lawnicons(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="14" cy="14" r="8.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="34" cy="34" r="8.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><rect width="17" height="17" x="5.5" y="25.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.309 6.5l-6.456 11.182a2 2 0 0 0 1.732 3h12.912a2 2 0 0 0 1.732-3L35.773 6.5a2 2 0 0 0-3.464 0Z"/>`),
		g.Group(children),
	)
}