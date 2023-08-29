package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowFatLineRightThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m234.83 125.17l-96-96A4 4 0 0 0 132 32v44H72a4 4 0 0 0-4 4v96a4 4 0 0 0 4 4h60v44a4 4 0 0 0 2.47 3.7a4 4 0 0 0 4.36-.87l96-96a4 4 0 0 0 0-5.66ZM140 214.34V176a4 4 0 0 0-4-4H76V84h60a4 4 0 0 0 4-4V41.66L226.34 128ZM44 80v96a4 4 0 0 1-8 0V80a4 4 0 0 1 8 0Z"/>`),
		g.Group(children),
	)
}