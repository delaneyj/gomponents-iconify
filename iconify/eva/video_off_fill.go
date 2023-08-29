package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaVideoOffFill0"><g id="evaVideoOffFill1"><path id="evaVideoOffFill2" fill="currentColor" d="M14.22 17.05L4.88 7.71L3.12 6L3 5.8A3 3 0 0 0 2 8v8a3 3 0 0 0 3 3h9a2.94 2.94 0 0 0 1.66-.51ZM21 7.15a1.7 1.7 0 0 0-1.85.3l-2.15 2V8a3 3 0 0 0-3-3H7.83l1.29 1.29l6.59 6.59l2 2l2 2a1.73 1.73 0 0 0 .6.11a1.68 1.68 0 0 0 .69-.15a1.6 1.6 0 0 0 1-1.48V8.63a1.6 1.6 0 0 0-1-1.48Zm-4 8.44l-2-2L8.41 7l-2-2l-1.7-1.71a1 1 0 0 0-1.42 1.42l.54.53L5.59 7l9.34 9.34l1.46 1.46l2.9 2.91a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42Z"/></g></g>`),
		g.Group(children),
	)
}