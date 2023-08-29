package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedalSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1111 512H937L580 0h888l-357 512zM1792 0q53 0 99 20t82 55t55 81t20 100q0 67-19 121t-57 109l-424 635q-91-114-217-179t-273-74L1663 0h129zM990 868q-146 8-272 73t-218 180L75 484q-36-54-55-108T0 256q0-53 20-99t55-82t81-55T256 0h129l605 868zm34 156q106 0 199 40t163 109t110 163t40 200q0 106-40 199t-109 163t-163 110t-200 40q-106 0-199-40t-163-109t-110-163t-40-200q0-106 40-199t109-163t163-110t200-40z"/>`),
		g.Group(children),
	)
}