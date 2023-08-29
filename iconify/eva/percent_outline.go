package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PercentOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaPercentOutline0"><g id="evaPercentOutline1"><path id="evaPercentOutline2" fill="currentColor" d="M8 11a3.5 3.5 0 1 0-3.5-3.5A3.5 3.5 0 0 0 8 11Zm0-5a1.5 1.5 0 1 1-1.5 1.5A1.5 1.5 0 0 1 8 6Zm8 8a3.5 3.5 0 1 0 3.5 3.5A3.5 3.5 0 0 0 16 14Zm0 5a1.5 1.5 0 1 1 1.5-1.5A1.5 1.5 0 0 1 16 19Zm3.74-14.74a.89.89 0 0 0-1.26 0L4.26 18.48a.91.91 0 0 0-.26.63a.89.89 0 0 0 1.52.63L19.74 5.52a.89.89 0 0 0 0-1.26Z"/></g></g>`),
		g.Group(children),
	)
}