package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleArrowUpLeftFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M17 3.34a10 10 0 1 1-14.995 8.984L2 12l.005-.324A10 10 0 0 1 17 3.34zM15 8H9l-.117.007l-.149.029l-.105.035l-.113.054l-.111.071a1.01 1.01 0 0 0-.112.097l-.08.09l-.067.096l-.052.098l-.044.11l-.03.112l-.017.126L8 15l.007.117A1 1 0 0 0 9 16l.117-.007A1 1 0 0 0 10 15v-3.585l4.293 4.292l.094.083a1 1 0 0 0 1.32-1.497L11.415 10H15l.117-.007A1 1 0 0 0 15 8z"/></g>`),
		g.Group(children),
	)
}