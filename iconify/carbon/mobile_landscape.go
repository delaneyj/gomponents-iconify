package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileLandscape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3 10v12a2.002 2.002 0 0 0 2 2h22a2.002 2.002 0 0 0 2-2V10a2.003 2.003 0 0 0-2-2H5a2.002 2.002 0 0 0-2 2Zm2 0h2v12H5Zm22 12H9V10h18Z"/>`),
		g.Group(children),
	)
}