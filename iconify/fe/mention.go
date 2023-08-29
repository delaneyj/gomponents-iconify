package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mention(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feMention0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feMention1" fill="currentColor" fill-rule="nonzero"><path id="feMention2" d="M21 12a3.5 3.5 0 0 1-5.909 2.539A4 4 0 1 1 14 8.535V8h2v4a1.5 1.5 0 0 0 3 0a7 7 0 1 0-3.392 6h3.1A9 9 0 1 1 21 12Zm-9 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></g>`),
		g.Group(children),
	)
}