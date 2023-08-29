package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileAudioSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 3v26h20V3zm2 2h16v22H8zm8 4.719v6.469A2.95 2.95 0 0 0 15 16c-1.645 0-3 1.355-3 3s1.355 3 3 3s3-1.355 3-3v-6.719l2.75.688l.5-1.938zM15 18c.563 0 1 .438 1 1c0 .563-.438 1-1 1c-.563 0-1-.438-1-1c0-.563.438-1 1-1z"/>`),
		g.Group(children),
	)
}