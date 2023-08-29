package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoCameraOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.696 5.078L5.398 6H4.5a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h11a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3h-.899l-.297-.922A3 3 0 0 0 11.449 3H8.55a3 3 0 0 0-2.855 2.078ZM4.5 8h2.354L7.6 5.693A1 1 0 0 1 8.551 5h2.898a1 1 0 0 1 .951.693L13.145 8H15.5a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6.5 11.5a3.5 3.5 0 1 0 7 0a3.5 3.5 0 0 0-7 0Zm5 0a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}