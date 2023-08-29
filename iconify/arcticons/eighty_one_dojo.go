package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EightyOneDojo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.546 15.787a3.25 3.25 0 0 0-3.25 3.25h0a3.25 3.25 0 0 0 3.25 3.25h2.113a3.25 3.25 0 0 0 3.25-3.25h0a3.25 3.25 0 0 0-3.25-3.25"/><path d="M14.659 15.787a3.25 3.25 0 0 0 3.25-3.25h0a3.25 3.25 0 0 0-3.25-3.25h-2.113a3.25 3.25 0 0 0-3.25 3.25h0a3.25 3.25 0 0 0 3.25 3.25m0 0h2.113"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21.159 11.057l3.25-1.77v13m-7.25 13.063v-8h1.8a3.5 3.5 0 0 1 3.5 3.5v1a3.5 3.5 0 0 1-3.5 3.5h-1.8Zm9.3.3h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v1.3a2 2 0 0 1-2 2Zm10 0h0a2 2 0 0 1-2-2v-1.3a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v1.3a2 2 0 0 1-2 2Zm-4.868-5.15v6c0 1.105-.887 2-1.981 2h0a1.966 1.966 0 0 1-1.401-.586"/><circle cx="31.591" cy="28.1" r=".75" fill="currentColor"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/>`),
		g.Group(children),
	)
}