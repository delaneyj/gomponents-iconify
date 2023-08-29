package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BdoDigitalBanking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33c-1.1 0-2 .9-2 2v33c0 1.1.9 2 2 2h33c1.1 0 2-.9 2-2v-33c0-1.1-.9-2-2-2Z"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M20.595 29.413V18.587h2.436a4.736 4.736 0 0 1 4.736 4.736v1.354a4.736 4.736 0 0 1-4.737 4.736h-2.435ZM14.966 24a2.707 2.707 0 0 1 0 5.413H10.5V18.587h4.466a2.707 2.707 0 0 1 0 5.413h0Zm0 0H10.5"/><rect width="7.172" height="10.826" x="30.328" y="18.587" rx="3.586" ry="3.586"/></g>`),
		g.Group(children),
	)
}