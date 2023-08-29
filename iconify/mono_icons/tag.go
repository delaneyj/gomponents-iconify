package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M2 3a1 1 0 0 1 1-1h8a1 1 0 0 1 .707.293l10 10a1 1 0 0 1 0 1.414l-8 8a1 1 0 0 1-1.414 0l-10-10A1 1 0 0 1 2 11V3zm2 1v6.586l9 9L19.586 13l-9-9H4z"/><path d="M9 7.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0z"/></g>`),
		g.Group(children),
	)
}