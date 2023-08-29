package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaMenuOutline0"><g id="evaMenuOutline1"><g id="evaMenuOutline2" fill="currentColor"><rect width="18" height="2" x="3" y="11" rx=".95" ry=".95"/><rect width="18" height="2" x="3" y="16" rx=".95" ry=".95"/><rect width="18" height="2" x="3" y="6" rx=".95" ry=".95"/></g></g></g>`),
		g.Group(children),
	)
}