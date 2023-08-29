package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 42H48a22 22 0 0 0-22 22v112a22 22 0 0 0 22 22h74v20H96a6 6 0 0 0 0 12h64a6 6 0 0 0 0-12h-26v-20h74a22 22 0 0 0 22-22V64a22 22 0 0 0-22-22ZM48 54h160a10 10 0 0 1 10 10v82H38V64a10 10 0 0 1 10-10Zm160 132H48a10 10 0 0 1-10-10v-18h180v18a10 10 0 0 1-10 10Z"/>`),
		g.Group(children),
	)
}