package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MidiInput(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2zm8.182 10c0-3.82-2.635-7.039-6.182-7.934V6h-4V4.066C6.453 4.96 3.818 8.18 3.818 12c0 4.51 3.673 8.182 8.182 8.182c4.51 0 8.182-3.673 8.182-8.182zM7 10.636a1.364 1.364 0 1 1 0 2.728a1.364 1.364 0 0 1 0-2.728zm10 0a1.364 1.364 0 1 1 0 2.728a1.364 1.364 0 0 1 0-2.728zm-8.636 3.637a1.364 1.364 0 1 1 0 2.727a1.364 1.364 0 0 1 0-2.727zm7.272 0a1.364 1.364 0 1 1 0 2.727a1.364 1.364 0 0 1 0-2.727zM12 15.636a1.364 1.364 0 1 1 0 2.728a1.364 1.364 0 0 1 0-2.728z" fill="currentColor"/>`),
		g.Group(children),
	)
}