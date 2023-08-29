package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlacrittyAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M475.429 36.571H36.57C16.371 36.571 0 52.943 0 73.143v365.714c0 20.2 16.371 36.572 36.571 36.572H475.43c20.2 0 36.571-16.372 36.571-36.572V73.143c0-20.2-16.371-36.572-36.571-36.572zM288.227 324.696L256 444.907l-32.752-120.211L256 245.976l32.227 78.72zm65.688 61.544L256 156.09l-97.915 230.15H107.17L228.584 84.907h54.832L404.831 386.24h-50.916z"/>`),
		g.Group(children),
	)
}