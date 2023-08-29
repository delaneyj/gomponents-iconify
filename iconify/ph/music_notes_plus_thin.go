package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNotesPlusThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M228 56a4 4 0 0 1-4 4h-20v20a4 4 0 0 1-8 0V60h-20a4 4 0 0 1 0-8h20V32a4 4 0 0 1 8 0v20h20a4 4 0 0 1 4 4ZM84 115.12V204a32.06 32.06 0 1 1-8-21.13V64a4 4 0 0 1 3-3.88l56-14a4 4 0 0 1 2 7.76L84 67.12v39.76l75-18.76a4 4 0 0 1 2 7.76ZM76 204a24 24 0 1 0-24 24a24 24 0 0 0 24-24Zm136-84v52a32.06 32.06 0 1 1-8-21.13V120a4 4 0 0 1 8 0Zm-8 52a24 24 0 1 0-24 24a24 24 0 0 0 24-24Z"/>`),
		g.Group(children),
	)
}