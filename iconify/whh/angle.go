package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.06 1024h-896q-26 0-45-18.5T.06 960V64q0-26 19-45t45.5-19t45 19t18.5 45v832h832q27 0 45.5 19t18.5 45.5t-18.5 45t-45.5 18.5zm-768-885V74q194 28 357 133t268 268t133 357h-65q-27-176-123.5-324.5t-245-245T192.06 139z"/>`),
		g.Group(children),
	)
}