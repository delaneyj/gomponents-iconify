package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CertificateTwoOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 12a3 3 0 1 0 3 3m-4-8h3"/><path d="M10 18v4l2-1l2 1v-4m-4 1H8a2 2 0 0 1-2-2V6m1.18-2.825C7.43 3.063 7.709 3 8 3h8a2 2 0 0 1 2 2v9m-.175 3.82A2 2 0 0 1 16 19h-2M3 3l18 18"/></g>`),
		g.Group(children),
	)
}