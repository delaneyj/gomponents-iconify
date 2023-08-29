package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fortnite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.797 41.478V6.522H30.91l-.644 7.67h-5.973v5.739h5.095v7.377H24.41v12.766l-7.613 1.404Zm0-3.366L4.5 39.839l1.113-21.9l-.996-9.485h12.18"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.748 8.454h12.693l-1.112 8.541L43.5 38.024L24.41 36.92"/>`),
		g.Group(children),
	)
}