package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaVideoOffOutline0"><g id="evaVideoOffOutline1"><path id="evaVideoOffOutline2" fill="currentColor" d="m17 15.59l-2-2L8.41 7l-2-2l-1.7-1.71a1 1 0 0 0-1.42 1.42l.54.53L5.59 7l9.34 9.34l1.46 1.46l2.9 2.91a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM14 17H5a1 1 0 0 1-1-1V8a1 1 0 0 1 .4-.78L3 5.8A3 3 0 0 0 2 8v8a3 3 0 0 0 3 3h9a2.94 2.94 0 0 0 1.66-.51L14.14 17a.7.7 0 0 1-.14 0Zm7-9.85a1.7 1.7 0 0 0-1.85.3l-2.15 2V8a3 3 0 0 0-3-3H7.83l2 2H14a1 1 0 0 1 1 1v4.17l4.72 4.72a1.73 1.73 0 0 0 .6.11a1.68 1.68 0 0 0 .69-.15a1.6 1.6 0 0 0 1-1.48V8.63A1.6 1.6 0 0 0 21 7.15Zm-1 7.45L17.19 12L20 9.4Z"/></g></g>`),
		g.Group(children),
	)
}