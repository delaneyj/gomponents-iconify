package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForbiddenCircleBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M4.929 4.929c-3.905 3.905-3.905 10.237 0 14.142c3.905 3.905 10.237 3.905 14.142 0c3.905-3.905 3.905-10.237 0-14.142c-3.905-3.905-10.237-3.905-14.142 0Z" opacity=".5"/><path d="M18.521 4.418L4.418 18.521a10.127 10.127 0 0 0 1.06 1.061L19.583 5.479a10.105 10.105 0 0 0-1.06-1.06Z"/></g>`),
		g.Group(children),
	)
}