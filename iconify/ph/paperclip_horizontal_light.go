package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperclipHorizontalLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M246 128a54.06 54.06 0 0 1-54 54H48a38 38 0 0 1 0-76h144a22 22 0 0 1 0 44H80a6 6 0 0 1 0-12h112a10 10 0 0 0 0-20H48a26 26 0 0 0 0 52h144a42 42 0 0 0 0-84H80a6 6 0 0 1 0-12h112a54.06 54.06 0 0 1 54 54Z"/>`),
		g.Group(children),
	)
}