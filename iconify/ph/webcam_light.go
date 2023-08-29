package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebcamLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M166 104a38 38 0 1 0-38 38a38 38 0 0 0 38-38Zm-64 0a26 26 0 1 1 26 26a26 26 0 0 1-26-26Zm122 98h-90v-20.25a78 78 0 1 0-12 0V202H32a6 6 0 0 0 0 12h192a6 6 0 0 0 0-12ZM62 104a66 66 0 1 1 66 66a66.08 66.08 0 0 1-66-66Z"/>`),
		g.Group(children),
	)
}