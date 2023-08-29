package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceTabletLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 26H64a22 22 0 0 0-22 22v160a22 22 0 0 0 22 22h128a22 22 0 0 0 22-22V48a22 22 0 0 0-22-22ZM54 70h148v116H54Zm10-32h128a10 10 0 0 1 10 10v10H54V48a10 10 0 0 1 10-10Zm128 180H64a10 10 0 0 1-10-10v-10h148v10a10 10 0 0 1-10 10Z"/>`),
		g.Group(children),
	)
}