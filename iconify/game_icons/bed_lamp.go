package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BedLamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m165.563 25l-87 174H432.93L339.207 25H165.563zM233 217v30h46v-30h-46zm-9 48c-1 0-9.308 1.608-18.52 5.15c-9.21 3.543-20.243 8.823-30.648 15.444C154.023 298.836 137 317 137 336c0 59.297 28.834 104.436 59.836 151h118.328C346.166 440.436 375 395.297 375 336c0-19-17.023-37.164-37.832-50.406c-10.405-6.621-21.437-11.9-30.648-15.444C297.308 266.608 289 265 288 265h-64z"/>`),
		g.Group(children),
	)
}