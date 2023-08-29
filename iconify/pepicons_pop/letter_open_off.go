package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterOpenOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.5 8.5v9h-15v-9"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m10 14.5l-6.5-4.33V5.5h13v4.67L10 14.5Zm-4.5-9l4.5-3l4.5 3"/><path stroke="currentColor" stroke-linecap="round" d="M7 8h6m-5.5 2h5"/><path fill="currentColor" fill-rule="evenodd" d="M2.5 7.5a1 1 0 0 1 1 1v8h13v-8a1 1 0 1 1 2 0v9a1 1 0 0 1-1 1h-15a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M2.5 5.5a1 1 0 0 1 1-1h13a1 1 0 0 1 1 1v4.67a1 1 0 0 1-.446.832l-6.5 4.33a1 1 0 0 1-1.108 0l-6.5-4.33a1 1 0 0 1-.446-.832V5.5Zm2 1v3.135l5.5 3.663l5.5-3.663V6.5h-11Z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M9.445 1.668a1 1 0 0 1 1.11 0l4.5 3l-1.11 1.664L10 3.702l-3.945 2.63l-1.11-1.664l4.5-3ZM6.5 8a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 0 1H7a.5.5 0 0 1-.5-.5Zm.5 2a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 0 1h-5A.5.5 0 0 1 7 10Z" clip-rule="evenodd"/><path fill="currentColor" d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}