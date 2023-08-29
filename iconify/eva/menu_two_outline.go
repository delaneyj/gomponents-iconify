package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaMenu2Outline0"><g id="evaMenu2Outline1"><g id="evaMenu2Outline2" fill="currentColor"><circle cx="4" cy="12" r="1"/><rect width="14" height="2" x="7" y="11" rx=".94" ry=".94"/><rect width="18" height="2" x="3" y="16" rx=".94" ry=".94"/><rect width="18" height="2" x="3" y="6" rx=".94" ry=".94"/></g></g></g>`),
		g.Group(children),
	)
}