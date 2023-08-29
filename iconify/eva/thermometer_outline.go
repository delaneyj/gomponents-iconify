package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaThermometerOutline0"><g id="evaThermometerOutline1"><path id="evaThermometerOutline2" fill="currentColor" d="M12 22a5 5 0 0 1-3-9V5a3 3 0 0 1 3-3a3 3 0 0 1 3 3v8a5 5 0 0 1-3 9Zm0-18a1 1 0 0 0-1 1v8.54a1 1 0 0 1-.5.87A3 3 0 0 0 9 17a3 3 0 0 0 6 0a3 3 0 0 0-1.5-2.59a1 1 0 0 1-.5-.87V5a.93.93 0 0 0-.29-.69A1 1 0 0 0 12 4Z"/></g></g>`),
		g.Group(children),
	)
}