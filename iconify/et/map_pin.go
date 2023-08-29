package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M12 0C5.383 0 0 5.394 0 12.022c0 9.927 11.201 19.459 11.678 19.86a.497.497 0 0 0 .64.004C12.795 31.492 24 22.124 24 12.022C24 5.394 18.617 0 12 0zm.002 30.838C10.161 29.193 1 20.579 1 12.022C1 5.944 5.935 1 12 1s11 4.944 11 11.022c0 8.702-9.152 17.193-10.998 18.816z"/><path d="M12 6c-3.309 0-6 2.691-6 6s2.691 6 6 6s6-2.691 6-6s-2.691-6-6-6zm0 11c-2.757 0-5-2.243-5-5s2.243-5 5-5s5 2.243 5 5s-2.243 5-5 5z"/></g>`),
		g.Group(children),
	)
}