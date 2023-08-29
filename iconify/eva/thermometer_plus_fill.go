package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThermometerPlusFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaThermometerPlusFill0"><g id="evaThermometerPlusFill1"><g id="evaThermometerPlusFill2" fill="currentColor"><rect width="6" height="2" x="2" y="5" rx="1" ry="1"/><rect width="6" height="2" x="2" y="5" rx="1" ry="1" transform="rotate(-90 5 6)"/><path d="M14 22a5 5 0 0 1-3-9V5a3 3 0 0 1 3-3a3 3 0 0 1 3 3v8a5 5 0 0 1-3 9Zm1-12.46V5a.93.93 0 0 0-.29-.69A1 1 0 0 0 14 4a1 1 0 0 0-1 1v4.54Z"/></g></g></g>`),
		g.Group(children),
	)
}