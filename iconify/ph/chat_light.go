package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 50H40a14 14 0 0 0-14 14v160a13.88 13.88 0 0 0 8.09 12.69A14.11 14.11 0 0 0 40 238a13.87 13.87 0 0 0 9-3.31l.09-.08l32.14-28.17a2 2 0 0 1 1.27-.44H216a14 14 0 0 0 14-14V64a14 14 0 0 0-14-14Zm2 142a2 2 0 0 1-2 2H82.5a14 14 0 0 0-9 3.29l-.09.08l-32.16 28.17A2 2 0 0 1 38 224V64a2 2 0 0 1 2-2h176a2 2 0 0 1 2 2Z"/>`),
		g.Group(children),
	)
}