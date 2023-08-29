package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Becu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.78 5.951h32.44a2 2 0 0 1 2 2v22.89a11.208 11.208 0 0 1-11.207 11.208H7.78a2 2 0 0 1-2-2V7.951a2 2 0 0 1 2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.544 24a4.75 4.75 0 1 1 0 9.5h-7.838v-19h7.838a4.75 4.75 0 1 1 0 9.5Zm0 0h-7.838"/>`),
		g.Group(children),
	)
}