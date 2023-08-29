package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 4l16 16M6.5 10.5A5.5 5.5 0 0 0 12 16v4m-4 0h4m0 0h4M9.5 9V6.5a2.5 2.5 0 0 1 5 0v4a2.5 2.5 0 0 1-1.5 2.292m4.5-2.292c0 1.93-.803 3.523-2.309 4.504"/>`),
		g.Group(children),
	)
}