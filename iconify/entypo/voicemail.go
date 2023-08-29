package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voicemail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.4 5.801a4.6 4.6 0 0 0-3.795 7.2H8.394A4.6 4.6 0 1 0 4.6 15h10.8a4.6 4.6 0 0 0 0-9.199zM2 10.4a2.6 2.6 0 1 1 5.2 0a2.6 2.6 0 0 1-5.2 0zM15.4 13a2.6 2.6 0 1 1-.002-5.2A2.6 2.6 0 0 1 15.4 13z"/>`),
		g.Group(children),
	)
}