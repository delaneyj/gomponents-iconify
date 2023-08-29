package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kodiremote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="29.97" height="29.97" x="9.02" y="9.02" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" transform="rotate(45 23.999 24)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.83 37.02V15.16H31.8v7.71L16.74 37.93m17.84-3.32l-7.26-7.25"/>`),
		g.Group(children),
	)
}