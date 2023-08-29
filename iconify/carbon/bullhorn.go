package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bullhorn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 6v2.17l-20.36 3.7a2 2 0 0 0-1.64 2v4.34a2 2 0 0 0 1.64 2l2.36.35V24a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-1.26l6 1.09V26h2V6Zm-8 18h-8v-3.07l8 1.45ZM6 18.17v-4.34l20-3.63v11.6Z"/>`),
		g.Group(children),
	)
}