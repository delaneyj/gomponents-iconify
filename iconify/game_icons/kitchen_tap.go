package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KitchenTap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M305.3 67.89L73 150.4V199h46v-32.7l191.6-72.54l-5.3-25.87zM369.5 169L119 252.5V217H73v222h61.1c-2.5-34.5 8.2-65.6 26.3-92.2c20-29.2 48.5-53.3 78.6-72.7c30-19.4 61.6-34 87.8-43.9c13.1-4.9 24.9-8.7 34.5-11.2c9.7-2.6 16.8-4 22.7-4h55v-46h-69.5zm46.5 85.2c-4 5.2-9.4 12.3-14.5 20.4c-10.4 16.1-18.8 36.1-17.2 46.1c1.1 7.3 5.7 16.3 11.9 22.9c6.2 6.6 13.6 10.5 19.8 10.5c6.1 0 13.5-3.9 19.7-10.5c6.2-6.6 10.8-15.6 11.9-22.9c1.6-10-6.8-30-17.1-46.1c-5.2-8.1-10.5-15.2-14.5-20.4zM41 457v30h138.5l-10-30H41z"/>`),
		g.Group(children),
	)
}