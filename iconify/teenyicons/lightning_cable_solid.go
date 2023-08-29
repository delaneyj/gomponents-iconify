package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningCableSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 0A1.5 1.5 0 0 0 4 1.5V5h7V1.5A1.5 1.5 0 0 0 9.5 0h-4ZM6 2h3v1H6V2Z" clip-rule="evenodd"/><path fill="currentColor" d="M3 6h9v5.5a1.5 1.5 0 0 1-1.5 1.5H10v2H9v-2H6v2H5v-2h-.5A1.5 1.5 0 0 1 3 11.5V6Z"/>`),
		g.Group(children),
	)
}