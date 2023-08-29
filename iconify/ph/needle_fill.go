package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeedleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M212.28 43.72a40 40 0 0 0-56.56 0l-24 24a8 8 0 0 0-2.23 4.3c-8.8 51.26-93.49 136.71-95.13 138.31a8 8 0 0 0 11.31 11.32c.86-.87 86.83-86.31 138.32-95.15a8 8 0 0 0 4.3-2.23l24-24a40 40 0 0 0 0-56.56Zm-22.62 33.94l-16 16a8 8 0 0 1-11.32-11.32l16-16a8 8 0 0 1 11.32 11.32Z"/>`),
		g.Group(children),
	)
}