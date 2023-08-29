package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M6.426 4.503L14.52 4.5a1 1 0 0 1 .997.92l.894 10.999a1 1 0 0 1-.916 1.078l-.08.003H5.58a1 1 0 0 1-1-1l.003-.077l.846-10.997a1 1 0 0 1 .997-.923z"/><path d="M13.5 8.5v.645c0 1.105-1.895 1.355-3 1.355s-3-.395-3-1.5v-.5"/></g>`),
		g.Group(children),
	)
}