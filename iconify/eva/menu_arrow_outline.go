package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MenuArrowOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaMenuArrowOutline0"><g id="evaMenuArrowOutline1"><g id="evaMenuArrowOutline2" fill="currentColor"><path d="M20.05 11H5.91l1.3-1.29a1 1 0 0 0-1.42-1.42l-3 3a1 1 0 0 0 0 1.42l3 3a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42L5.91 13h14.14a1 1 0 0 0 .95-.95V12a1 1 0 0 0-.95-1Z"/><rect width="18" height="2" x="3" y="17" rx=".95" ry=".95"/><rect width="18" height="2" x="3" y="5" rx=".95" ry=".95"/></g></g></g>`),
		g.Group(children),
	)
}