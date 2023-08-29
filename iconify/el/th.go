package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Th(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 0v300h300V0H0zm450 0v300h300V0H450zm450 0v300h300V0H900zM0 450v300h300V450H0zm450 0v300h300V450H450zm450 0v300h300V450H900zM0 900v300h300V900H0zm450 0v300h300V900H450zm450 0v300h300V900H900z"/>`),
		g.Group(children),
	)
}