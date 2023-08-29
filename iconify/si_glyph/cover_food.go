package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoverFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M16.471 12.076H1.49c-.271 0-.489.207-.489.461c0 .256.219.463.489.463h14.981c.271 0 .49-.207.49-.463c0-.254-.22-.461-.49-.461zM8.987 4.201c-3.839 0-6.95 3.008-6.95 6.718h13.9c.001-3.71-3.112-6.718-6.95-6.718z"/><path d="M8 3h1.969v1.969H8z"/></g>`),
		g.Group(children),
	)
}