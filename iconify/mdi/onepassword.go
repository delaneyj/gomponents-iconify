package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Onepassword(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1C5.92 1 1 5.92 1 12s4.92 11 11 11s11-4.92 11-11S18.08 1 12 1m0 19a8 8 0 0 1-8-8a8 8 0 0 1 8-8a8 8 0 0 1 8 8a8 8 0 0 1-8 8m1-6.5c0 .63.4 1.2 1 1.41V18h-4v-6.09c.78-.27 1.19-1.11.93-1.91a1.5 1.5 0 0 0-.93-.91V6h4v6.09c-.6.21-1 .78-1 1.41Z"/>`),
		g.Group(children),
	)
}