package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveYOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd" opacity=".2"><path d="M6.348 8.294A1.5 1.5 0 0 1 6.54 6.18l4-3.333a1.5 1.5 0 0 1 1.92 2.304l-4 3.334a1.5 1.5 0 0 1-2.112-.192Z"/><path d="M16.652 8.294a1.5 1.5 0 0 1-2.112.192l-4-3.334a1.5 1.5 0 0 1 1.92-2.304l4 3.333a1.5 1.5 0 0 1 .192 2.113Z"/><path d="M11.5 4.5A1.5 1.5 0 0 1 13 6v8a1.5 1.5 0 0 1-3 0V6a1.5 1.5 0 0 1 1.5-1.5Z"/><path d="M16.652 13.707a1.5 1.5 0 0 1-.192 2.112l-4 3.333a1.5 1.5 0 0 1-1.92-2.304l4-3.333a1.5 1.5 0 0 1 2.112.191Z"/><path d="M6.348 13.707a1.5 1.5 0 0 1 2.112-.193l4 3.334a1.5 1.5 0 1 1-1.92 2.305l-4-3.334a1.5 1.5 0 0 1-.192-2.113Z"/><path d="M11.5 17.5A1.5 1.5 0 0 1 10 16V8a1.5 1.5 0 0 1 3 0v8a1.5 1.5 0 0 1-1.5 1.5Z"/></g><path fill-rule="evenodd" d="M5.616 6.653a.5.5 0 0 1 .064-.704l4-3.333a.5.5 0 1 1 .64.768l-4 3.333a.5.5 0 0 1-.704-.064Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M14.384 6.653a.5.5 0 0 1-.704.064l-4-3.333a.5.5 0 1 1 .64-.768l4 3.333a.5.5 0 0 1 .064.704ZM10 4.5a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5Zm4.384 8.847a.5.5 0 0 1-.064.704l-4 3.333a.5.5 0 0 1-.64-.768l4-3.333a.5.5 0 0 1 .704.064Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5.616 13.347a.5.5 0 0 1 .704-.064l4 3.333a.5.5 0 0 1-.64.768l-4-3.333a.5.5 0 0 1-.064-.704Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M10 17a.5.5 0 0 1-.5-.5v-13a.5.5 0 0 1 1 0v13a.5.5 0 0 1-.5.5Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}