package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.585 192h-2l2 444v4q0 26-19 45t-45 19h-768q-27 0-45.5-19t-18.5-45v-4l2-444h-2q-27 0-45.5-18.5T.585 128t18.5-45.5t45-18.5t45.5 18.5t19 45.5q0 33-27 52l219 310l166-368q-17-7-27.5-23t-10.5-35q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45.5q0 19-10.5 35t-27.5 23l166 368l219-310q-27-19-27-52q0-27 18.5-45.5t45-18.5t45.5 18.5t19 45.5t-19 45.5t-45 18.5zm-896 576h896q26 0 45 18.5t19 45t-19 45.5t-45 19h-896q-26 0-45-19t-19-45.5t18.5-45t45.5-18.5z"/>`),
		g.Group(children),
	)
}