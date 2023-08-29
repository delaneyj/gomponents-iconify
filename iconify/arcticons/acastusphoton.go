package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Acastusphoton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.106 9.606a6.38 6.38 0 0 0-6.38 6.38c0 4.992 4.884 11.01 6.148 12.485a.385.385 0 0 0 .589-.003c1.243-1.48 6.022-7.493 6.022-12.482a6.38 6.38 0 0 0-6.38-6.38Zm0 8.764a2.384 2.384 0 1 1 2.384-2.385a2.384 2.384 0 0 1-2.384 2.385Z"/><circle cx="19" cy="19" r="13.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 42.5L28.541 28.541"/>`),
		g.Group(children),
	)
}