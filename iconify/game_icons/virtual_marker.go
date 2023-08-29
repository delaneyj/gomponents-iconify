package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VirtualMarker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 32L32 224h128l96-80l96 80h128L256 32zm0 176l-52 39H32v18h172l52 39l52-39h172v-18H308l-52-39zM32 288l224 192l224-192H352l-96 80l-96-80H32z"/>`),
		g.Group(children),
	)
}