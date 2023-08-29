package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Overhead(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M37.368 144.56c97.187 0 101.127 1.314 101.127 1.314C263.844-57.842 363.082 219.342 375.947 319.62l-36.72.16l71.834 125.393l68.982-125.01l-35.724-1.472c-82.41-287.54-303.28-300.208-406.95-174.13z"/>`),
		g.Group(children),
	)
}