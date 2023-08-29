package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChainedArrowHeads(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M97.53 20.705v.002l6.425 82.932l-86.434-6.427v53.73l207.912 74.754L151.26 20.705H97.53zm-80.01.002v31.537l54.9 19.858l-18.69-51.395H17.52zM237.366 45.21l15.754 207.743L45.374 237.2l300.363 108.374L237.368 45.21h-.003zm117.342 171.927l.002.008v-.008h-.003zm.002.008l11.272 148.67l-148.68-11.272l214.968 77.562l-77.56-214.96zm87.493 137.65l.002.008v-.008h-.002zm.002.008l7.158 94.396l-94.404-7.16l136.49 49.247l-49.245-136.484z"/>`),
		g.Group(children),
	)
}