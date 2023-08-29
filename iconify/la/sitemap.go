package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sitemap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 5v8h3v2H5v4H2v8h8v-8H7v-2h8v2h-3v8h8v-8h-3v-2h8v2h-3v8h8v-8h-3v-4H17v-2h3V5zm2 2h4v4h-4zM4 21h4v4H4zm10 0h4v4h-4zm10 0h4v4h-4z"/>`),
		g.Group(children),
	)
}