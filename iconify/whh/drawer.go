package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drawer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.417 1024h-768q-26 0-45-18.5t-19-45.5V64q0-26 19-45t45-19h768q27 0 45.5 19t18.5 45v896q0 27-18.5 45.5t-45.5 18.5zm-64-864q0-13-9.5-22.5t-22.5-9.5h-576q-13 0-22.5 9.5t-9.5 22.5v256q0 13 9.5 22.5t22.5 9.5h576q13 0 22.5-9.5t9.5-22.5V160zm0 448q0-13-9.5-22.5t-22.5-9.5h-576q-13 0-22.5 9.5t-9.5 22.5v256q0 13 9.5 22.5t22.5 9.5h576q13 0 22.5-9.5t9.5-22.5V608zm-320.5 224q-26.5 0-45-19t-18.5-45t18.5-45t45-19t45.5 19t19 45t-19 45t-45.5 19zm0-448q-26.5 0-45-18.5t-18.5-45.5t18.5-45.5t45-18.5t45.5 19t19 45.5t-19 45t-45.5 18.5z"/>`),
		g.Group(children),
	)
}