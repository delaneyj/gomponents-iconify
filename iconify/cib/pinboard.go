package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pinboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m17.802 19.448l-6.01 6.151l.958-5.417l-8.177-10.089L0 10.338L10.104-.001v3.938l9.62 8.859l6.021-1.474l-6.255 6.646L32 31.999z"/>`),
		g.Group(children),
	)
}