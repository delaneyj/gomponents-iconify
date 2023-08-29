package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControllerOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M12.25 10a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Zm2 2.5a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Z"/><path fill-rule="evenodd" d="M14.5 4h-9a5 5 0 0 0-5 5v2a5 5 0 0 0 5 5h9a5 5 0 0 0 5-5V9a5 5 0 0 0-5-5Zm-12 5a3 3 0 0 1 3-3h9a3 3 0 0 1 3 3v2a3 3 0 0 1-3 3h-9a3 3 0 0 1-3-3V9Z" clip-rule="evenodd"/><path d="M4 11a1 1 0 1 1 0-2h4a1 1 0 0 1 0 2H4Z"/><path d="M7 12a1 1 0 1 1-2 0V8a1 1 0 0 1 2 0v4Z"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}