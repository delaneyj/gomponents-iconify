package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Razor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92d3f5" d="M32.5 22.5s0 38 1 42s4 4 5 0c2.03-8.123 1-42 1-42m-17-7s10 5 10 7h7c0-1.699 7.214-5.562 9.388-6.687c.385-.2.612-.313.612-.313Z"/><rect width="32" height="9" x="20" y="6.5" fill="#d0cfce" rx="2.244" ry="2.244"/><g fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2"><path stroke-linecap="round" d="M32.5 22.5s0 38 1 42s4 4 5 0c2.03-8.123 1-42 1-42"/><rect width="32" height="9" x="20" y="6.5" stroke-linecap="round" rx="2.244" ry="2.244"/><path stroke-linecap="round" d="M22.5 15.5s10 5 10 7m7 0c0-1.699 7.214-5.562 9.388-6.687c.385-.2.612-.313.612-.313"/><path d="M20 11h32"/></g>`),
		g.Group(children),
	)
}