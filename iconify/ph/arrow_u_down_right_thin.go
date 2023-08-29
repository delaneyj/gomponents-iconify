package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUDownRightThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m226.83 170.83l-48 48a4 4 0 0 1-5.66-5.66L214.34 172H88a60 60 0 0 1 0-120h88a4 4 0 0 1 0 8H88a52 52 0 0 0 0 104h126.34l-41.17-41.17a4 4 0 0 1 5.66-5.66l48 48a4 4 0 0 1 0 5.66Z"/>`),
		g.Group(children),
	)
}