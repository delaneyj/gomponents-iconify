package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaThermometerFill0"><g id="evaThermometerFill1"><path id="evaThermometerFill2" fill="currentColor" d="M12 22a5 5 0 0 1-3-9V5a3 3 0 0 1 3-3a3 3 0 0 1 3 3v8a5 5 0 0 1-3 9Zm1-12.46V5a.93.93 0 0 0-.29-.69A1 1 0 0 0 12 4a1 1 0 0 0-1 1v4.54Z"/></g></g>`),
		g.Group(children),
	)
}