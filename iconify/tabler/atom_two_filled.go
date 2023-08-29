package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtomTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M12 8a4 4 0 1 1-3.995 4.2L8 12l.005-.2A4 4 0 0 1 12 8zm0 12a1 1 0 0 1 .993.883l.007.127a1 1 0 0 1-1.993.117L11 21a1 1 0 0 1 1-1zM3 8a1 1 0 0 1 .993.883L4 9.01a1 1 0 0 1-1.993.117L2 9a1 1 0 0 1 1-1zm18 0a1 1 0 0 1 .993.883L22 9.01a1 1 0 0 1-1.993.117L20 9a1 1 0 0 1 1-1zM2.89 12.006a1 1 0 0 1 1.104.884a8 8 0 0 0 4.444 6.311A1 1 0 1 1 7.562 21a10 10 0 0 1-5.556-7.89a1 1 0 0 1 .884-1.103zM20.993 12l.117.006a1 1 0 0 1 .884 1.104a10 10 0 0 1-5.556 7.889a1 1 0 1 1-.876-1.798a8 8 0 0 0 4.444-6.31a1 1 0 0 1 .987-.891zM5.567 4.226a10 10 0 0 1 12.666 0a1 1 0 1 1-1.266 1.548a8 8 0 0 0-10.134 0a1 1 0 1 1-1.266-1.548z"/></g>`),
		g.Group(children),
	)
}