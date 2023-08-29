package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zalando(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.27 24c-.88 0-1.36-.2-1.62-.36c-.36-.21-1.02-.75-1.62-2.33A27.06 27.06 0 0 1 .49 12c.02-3.66.59-6.76 1.54-9.3C2.63 1.1 3.29.56 3.65.35C3.91.21 4.39 0 5.27 0c.33 0 .72.03 1.18.1a26.1 26.1 0 0 1 8.7 3.3h.01a26.4 26.4 0 0 1 7.16 6.01c1.06 1.32 1.19 2.17 1.19 2.59c0 .42-.13 1.27-1.19 2.59a26.4 26.4 0 0 1-7.16 6h-.01a26.03 26.03 0 0 1-8.7 3.3c-.46.08-.85.11-1.18.11z"/>`),
		g.Group(children),
	)
}