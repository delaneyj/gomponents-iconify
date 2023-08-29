package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 0v300h300V0H0zm468.75 0v300H1200V0H468.75zM0 450v300h300V450H0zm468.75 0v300H1200V450H468.75zM0 900v300h300V900H0zm468.75 0v300H1200V900H468.75z"/>`),
		g.Group(children),
	)
}