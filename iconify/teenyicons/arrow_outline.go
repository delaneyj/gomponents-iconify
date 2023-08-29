package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M.5.5V0a.5.5 0 0 0-.5.5h.5Zm0 4H0a.5.5 0 0 0 .854.354L.5 4.5Zm4-4l.354.354A.5.5 0 0 0 4.5 0v.5ZM2.146 2.854l12 12l.708-.708l-12-12l-.708.708ZM0 .5v4h1v-4H0Zm.854 4.354l4-4l-.708-.708l-4 4l.708.708ZM4.5 0h-4v1h4V0Z"/>`),
		g.Group(children),
	)
}