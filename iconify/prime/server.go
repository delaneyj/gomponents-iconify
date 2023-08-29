package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Server(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 8.5v-3a1 1 0 0 0-1-1h-15a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h15a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1a1 1 0 0 0 1-1Zm-1 10h-15v-3h15Zm0-5h-15v-3h15Zm0-5h-15v-3h15Z"/><circle cx="6.25" cy="7" r=".75" fill="currentColor"/><circle cx="8.75" cy="7" r=".75" fill="currentColor"/><circle cx="6.25" cy="12" r=".75" fill="currentColor"/><circle cx="8.75" cy="12" r=".75" fill="currentColor"/><circle cx="6.25" cy="17" r=".75" fill="currentColor"/><circle cx="8.75" cy="17" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}