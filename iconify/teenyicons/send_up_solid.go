package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SendUpSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m7.5.793l4.354 4.353l-.708.708L8 2.707V12H7V2.707L3.854 5.854l-.708-.708L7.5.793ZM14 13v1H1v-1h13Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}