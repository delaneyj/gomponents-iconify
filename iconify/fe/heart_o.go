package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feHeartO0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feHeartO1" fill="currentColor"><path id="feHeartO2" d="M12 20c-2.205-.48-9-4.24-9-11a5 5 0 0 1 9-3a5 5 0 0 1 9 3c0 6.76-6.795 10.52-9 11Zm0-2c3.12-.93 7-4.805 7-9a3 3 0 0 0-3-3c-1.305 0-2.638.833-4 2.5C10.638 6.833 9.305 6 8 6a3 3 0 0 0-3 3c0 4.195 3.88 8.07 7 9Z"/></g></g>`),
		g.Group(children),
	)
}