package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LastTrackButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M52 46L34.6 32L52 18zm-17.4 0L17.2 32l17.4-14zM12 18h5.2v28H12z"/>`),
		g.Group(children),
	)
}