package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinnedFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1900 748q-41 40-93 61t-110 22q-44 0-84-12l-433 433q17 41 26 84t10 89q0 86-33 165t-93 140L318 958q61-60 140-93t165-33q45 0 88 9t85 27l433-433q-12-40-12-84q0-57 21-109t62-94l600 600z"/>`),
		g.Group(children),
	)
}