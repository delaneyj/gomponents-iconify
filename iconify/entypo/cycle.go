package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cycle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5.516 14.224c-2.262-2.432-2.222-6.244.128-8.611a6.074 6.074 0 0 1 3.414-1.736L8.989 1.8a8.112 8.112 0 0 0-4.797 2.351c-3.149 3.17-3.187 8.289-.123 11.531l-1.741 1.752l5.51.301l-.015-5.834l-2.307 2.323zm6.647-11.959l.015 5.834l2.307-2.322c2.262 2.434 2.222 6.246-.128 8.611a6.07 6.07 0 0 1-3.414 1.736l.069 2.076a8.122 8.122 0 0 0 4.798-2.35c3.148-3.172 3.186-8.291.122-11.531l1.741-1.754l-5.51-.3z"/>`),
		g.Group(children),
	)
}