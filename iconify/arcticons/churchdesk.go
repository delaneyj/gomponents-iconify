package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Churchdesk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.402 17.287V33.18a3.384 3.384 0 0 0 6.769 0h0c0-5.996-.152-6.432 0-21.761a6.912 6.912 0 0 0-13.824-.075v21.637a10.563 10.563 0 0 0 21.126 0V19.65m.335-8.93s1.845.98 1.845 3.965a2.228 2.228 0 1 1-4.449 0a3.56 3.56 0 0 1 .894-2.64a1.028 1.028 0 0 0 .978 1.123a1.012 1.012 0 0 0 .984-1.038v0a2.52 2.52 0 0 0-.252-1.41"/>`),
		g.Group(children),
	)
}