package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitkub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.753 24a2.778 2.778 0 1 1-.003-.142"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.188 35.672L4.517 24l11.671-11.672l8.163 8.163M29.247 24a2.778 2.778 0 1 1 .003.142"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.812 12.328L43.483 24l-11.67 11.672l-8.163-8.163"/>`),
		g.Group(children),
	)
}