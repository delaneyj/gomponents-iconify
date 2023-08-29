package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Repeat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.498 22a.997.997 0 0 1-.707-.293l-2.5-2.5a1 1 0 0 1 0-1.414l2.5-2.5a1 1 0 0 1 1.414 1.414L10.412 18.5l1.793 1.793A1 1 0 0 1 11.498 22z" opacity=".5"/><path fill="currentColor" d="M21 4.5h-2a1 1 0 0 0 0 2h1v11h-8.588l-1 1l1 1H21a1 1 0 0 0 1-1v-13a1 1 0 0 0-1-1z" opacity=".5"/><path fill="currentColor" d="M12.5 2c.265 0 .52.105.707.293l2.5 2.5a1 1 0 0 1 0 1.414l-2.5 2.5a1 1 0 0 1-1.414-1.414L13.586 5.5l-1.793-1.793A1 1 0 0 1 12.5 2z"/><path fill="currentColor" d="M5 17.5H4v-11h8.586l1-1l-1-1H3a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h2a1 1 0 0 0 0-2z"/>`),
		g.Group(children),
	)
}