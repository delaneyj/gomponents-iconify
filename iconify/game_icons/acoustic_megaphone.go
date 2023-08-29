package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AcousticMegaphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m443.535 120.186l-112 64l8.93 15.628l112-64l-8.93-15.628zM297 153v206h17.973V153H297zm-18 9.367L73 235.072v41.856l206 72.705V162.367zM39 240v32h18v-32H39zm297 7v18h128v-18H336zM99.332 300.89l-14.8 40.215L181.02 379.7l16.11-40.364l-16.716-6.672l-9.434 23.635l-63.51-25.405l8.755-23.786l-16.893-6.22zm241.133 11.296l-8.93 15.628l112 64l8.93-15.628l-112-64z"/>`),
		g.Group(children),
	)
}