package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MrOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#cd2a3e" d="M0 0h512v512H0z"/><path fill="#006233" d="M0 76.8h512v358.4H0z"/><path fill="#ffc400" d="M416 164.9a160 160 0 0 1-320 0a165.2 165.2 0 0 0-5.4 41.8A165.4 165.4 0 1 0 416 165z" class="mr-st1"/><path fill="#ffc400" d="m256 100l-14.4 44.3h-46.5l37.6 27.3l-14.3 44.2l37.6-27.3l37.6 27.3l-14.4-44.2l37.7-27.3h-46.5z"/>`),
		g.Group(children),
	)
}