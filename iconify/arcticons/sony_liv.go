package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SonyLiv(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 7.5v3.23a10 10 0 0 0 6.915 9.513l23.17 7.514A10 10 0 0 1 42.5 37.27v3.23"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M24.816 17v-7l4.638 7v-7m6.583 0l-2.319 3.5l-2.319-3.5m2.319 7v-3.5m-21.616 2.733c.429.559.967.767 1.716.767h1.036c.965 0 1.747-.782 1.747-1.746v-.008c0-.964-.782-1.746-1.747-1.746h-1.143a1.748 1.748 0 0 1-1.748-1.748h0c0-.968.784-1.752 1.752-1.752h1.03c.75 0 1.288.208 1.717.767"/><rect width="4.638" height="7" x="18.233" y="10" rx="2.319" ry="2.319"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.611 31.249L27.875 38.5l-2.736-7.251"/><circle cx="22.187" cy="27.897" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.187 31.249V38.5m-4.798-10.945v9.577c0 .756.612 1.368 1.368 1.368h.41"/>`),
		g.Group(children),
	)
}