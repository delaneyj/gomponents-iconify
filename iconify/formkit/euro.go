package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Euro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 15c-3.86 0-7-3.14-7-7s3.14-7 7-7s7 3.14 7 7s-3.14 7-7 7ZM8 2C4.69 2 2 4.69 2 8s2.69 6 6 6s6-2.69 6-6s-2.69-6-6-6Z"/><path fill="currentColor" d="M7.89 12c-1.35 0-2.44-1.1-2.44-2.44V6.45c0-1.35 1.1-2.44 2.44-2.44s2.44 1.1 2.44 2.44c0 .28-.22.5-.5.5s-.5-.22-.5-.5c0-.8-.65-1.44-1.44-1.44s-1.44.65-1.44 1.44v3.11c0 .8.65 1.44 1.44 1.44s1.44-.65 1.44-1.44c0-.28.22-.5.5-.5s.5.22.5.5c0 1.35-1.1 2.44-2.44 2.44Z"/><path fill="currentColor" d="M7.5 7.5H5.17c-.28 0-.5-.22-.5-.5s.22-.5.5-.5H7.5c.28 0 .5.22.5.5s-.22.5-.5.5Zm0 2H5.17c-.28 0-.5-.22-.5-.5s.22-.5.5-.5H7.5c.28 0 .5.22.5.5s-.22.5-.5.5Z"/>`),
		g.Group(children),
	)
}