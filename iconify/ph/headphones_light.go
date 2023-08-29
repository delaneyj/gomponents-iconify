package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphonesLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200.47 64.07A101.37 101.37 0 0 0 128.77 34H128A102 102 0 0 0 26 136v56a22 22 0 0 0 22 22h16a22 22 0 0 0 22-22v-40a22 22 0 0 0-22-22H38.2A90.12 90.12 0 0 1 192 72.52A89.41 89.41 0 0 1 217.81 130H192a22 22 0 0 0-22 22v40a22 22 0 0 0 22 22h16a22 22 0 0 0 22-22v-56a101.44 101.44 0 0 0-29.53-71.93ZM64 142a10 10 0 0 1 10 10v40a10 10 0 0 1-10 10H48a10 10 0 0 1-10-10v-50Zm154 50a10 10 0 0 1-10 10h-16a10 10 0 0 1-10-10v-40a10 10 0 0 1 10-10h26Z"/>`),
		g.Group(children),
	)
}