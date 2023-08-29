package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocExpand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 1.5v12h5.75a.75.75 0 0 1 0 1.5H3a1 1 0 0 1-1-1V1a1 1 0 0 1 1-1h6.644a1 1 0 0 1 .72.305l3.355 3.476a1 1 0 0 1 .281.695v.774a.75.75 0 0 1-.813.747a.76.76 0 0 1-.062.003H9.75A1.75 1.75 0 0 1 8 4.25V1.5H3.5Zm6 .07l2.828 2.93H9.75a.25.25 0 0 1-.25-.25V1.57Zm6.353 12.228L13.5 16l-2.353-2.202c-.314-.295-.091-.798.353-.798h1.25v-3H11.5c-.444 0-.667-.503-.353-.798L13.5 7l2.353 2.202c.314.295.091.798-.353.798h-1.25v3h1.25c.444 0 .667.503.353.798Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}