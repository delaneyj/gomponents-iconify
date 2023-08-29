package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unsplash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 21H3V10.875h5.625v5.063h6.75v-5.063H21V21ZM15.375 8.063h-6.75V3h6.75v5.063Z"/>`),
		g.Group(children),
	)
}