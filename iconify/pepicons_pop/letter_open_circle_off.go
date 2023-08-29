package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterOpenCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.5 11.5v9h-15v-9"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m13 17.5l-6.5-4.33V8.5h13v4.67L13 17.5Zm-4.5-9l4.5-3l4.5 3"/><path stroke="currentColor" stroke-linecap="round" d="M10 11h6m-5.5 2h5"/><path fill="currentColor" fill-rule="evenodd" d="M5.5 10.5a1 1 0 0 1 1 1v8h13v-8a1 1 0 1 1 2 0v9a1 1 0 0 1-1 1h-15a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M5.5 8.5a1 1 0 0 1 1-1h13a1 1 0 0 1 1 1v4.67a1 1 0 0 1-.446.832l-6.5 4.33a1 1 0 0 1-1.108 0l-6.5-4.33a1 1 0 0 1-.446-.832V8.5Zm2 1v3.135l5.5 3.663l5.5-3.663V9.5h-11Z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M12.445 4.668a1 1 0 0 1 1.11 0l4.5 3l-1.11 1.664L13 6.702l-3.945 2.63l-1.11-1.664l4.5-3ZM9.5 11a.5.5 0 0 1 .5-.5h6a.5.5 0 0 1 0 1h-6a.5.5 0 0 1-.5-.5Zm.5 2a.5.5 0 0 1 .5-.5h5a.5.5 0 0 1 0 1h-5a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill="currentColor" d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill="currentColor" fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}