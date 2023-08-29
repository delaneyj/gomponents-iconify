package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineageCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="27.238" cy="26.437" r="12.652" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.887 37.76h7.7a2.913 2.913 0 0 0 2.913-2.913v-15.04a4.574 4.574 0 0 0-2.928-4.268l-4.647-1.792c-.372-3.932-2.807-4.835-2.807-4.835H21.421s-2.435.903-2.807 4.834L8.147 15.914A4.574 4.574 0 0 0 4.5 20.392v14.455a2.913 2.913 0 0 0 2.913 2.913H21.59"/><circle cx="8.243" cy="19.821" r="1.373" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.581 15.203V37.76"/>`),
		g.Group(children),
	)
}