package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeAltOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 22H5a1 1 0 0 1-1-1v-8H2l9.292-9.293a1 1 0 0 1 1.415 0L22 13h-2v8a1 1 0 0 1-1 1Zm-9-7h4v5h4v-8.172l-6-6l-6 6V20h4v-5Z"/>`),
		g.Group(children),
	)
}