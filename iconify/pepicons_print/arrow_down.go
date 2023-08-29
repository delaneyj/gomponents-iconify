package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><g opacity=".2"><path d="M17.152 11.206a1.5 1.5 0 0 1-.192 2.113l-4 3.333a1.5 1.5 0 1 1-1.92-2.304l4-3.334a1.5 1.5 0 0 1 2.112.192Z"/><path d="M6.848 11.206a1.5 1.5 0 0 1 2.112-.192l4 3.334a1.5 1.5 0 1 1-1.92 2.304l-4-3.333a1.5 1.5 0 0 1-.192-2.113Z"/><path d="M12 15a1.5 1.5 0 0 1-1.5-1.5v-8a1.5 1.5 0 0 1 3 0v8A1.5 1.5 0 0 1 12 15Z"/></g><path d="M14.384 11.347a.5.5 0 0 1-.064.704l-4 3.333a.5.5 0 0 1-.64-.768l4-3.333a.5.5 0 0 1 .704.064Z"/><path d="M5.616 11.347a.5.5 0 0 1 .704-.064l4 3.333a.5.5 0 0 1-.64.768l-4-3.333a.5.5 0 0 1-.064-.704Z"/><path d="M10 15a.5.5 0 0 1-.5-.5V5a.5.5 0 0 1 1 0v9.5a.5.5 0 0 1-.5.5Z"/></g>`),
		g.Group(children),
	)
}