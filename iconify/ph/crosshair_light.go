package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrosshairLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 122h-10.2A94.13 94.13 0 0 0 134 34.2V24a6 6 0 0 0-12 0v10.2A94.13 94.13 0 0 0 34.2 122H24a6 6 0 0 0 0 12h10.2a94.13 94.13 0 0 0 87.8 87.8V232a6 6 0 0 0 12 0v-10.2a94.13 94.13 0 0 0 87.8-87.8H232a6 6 0 0 0 0-12Zm-98 87.76V200a6 6 0 0 0-12 0v9.76A82.09 82.09 0 0 1 46.24 134H56a6 6 0 0 0 0-12h-9.76A82.09 82.09 0 0 1 122 46.24V56a6 6 0 0 0 12 0v-9.76A82.09 82.09 0 0 1 209.76 122H200a6 6 0 0 0 0 12h9.76A82.09 82.09 0 0 1 134 209.76ZM128 90a38 38 0 1 0 38 38a38 38 0 0 0-38-38Zm0 64a26 26 0 1 1 26-26a26 26 0 0 1-26 26Z"/>`),
		g.Group(children),
	)
}