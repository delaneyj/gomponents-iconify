package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextTrackButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="m12 18l17.4 14L12 46zm17.4 0l17.4 14l-17.4 14zm17.4 0H52v28h-5.2z"/>`),
		g.Group(children),
	)
}