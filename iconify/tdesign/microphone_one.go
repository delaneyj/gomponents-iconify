package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 6a5 5 0 0 1 10 0v5a5 5 0 0 1-10 0V6Zm5-3a3 3 0 0 0-3 3v5a3 3 0 1 0 6 0V6a3 3 0 0 0-3-3Zm-6 7v1a6 6 0 0 0 12 0v-1h2v1a8.001 8.001 0 0 1-7 7.938V20h5v2H6v-2h5v-1.062A8.001 8.001 0 0 1 4 11v-1h2Z"/>`),
		g.Group(children),
	)
}