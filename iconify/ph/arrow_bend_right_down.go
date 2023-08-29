package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBendRightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m205.66 181.66l-48 48a8 8 0 0 1-11.32 0l-48-48a8 8 0 0 1 11.32-11.32L144 204.69V128a88.1 88.1 0 0 0-88-88a8 8 0 0 1 0-16a104.11 104.11 0 0 1 104 104v76.69l34.34-34.35a8 8 0 0 1 11.32 11.32Z"/>`),
		g.Group(children),
	)
}