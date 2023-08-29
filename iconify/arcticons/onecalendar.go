package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Onecalendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.447 5.5A1.943 1.943 0 0 0 5.5 7.447v33.106A1.943 1.943 0 0 0 7.447 42.5h33.106a1.943 1.943 0 0 0 1.947-1.947V7.447A1.943 1.943 0 0 0 40.553 5.5ZM5.5 12.645h37m-3.5 1.5v-3m-10 3v-3m-10 3v-3m-10 3v-3"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".961" d="M20.375 25.477L24 23.503m0 0v14.501m0 0a10.102 10.102 0 1 0-3.989-.817"/>`),
		g.Group(children),
	)
}