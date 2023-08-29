package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuestionCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2zm0 7a1 1 0 0 0-1 1a1 1 0 1 1-2 0a3 3 0 1 1 4.44 2.633a1.404 1.404 0 0 0-.383.288a.303.303 0 0 0-.057.085v.494a1 1 0 1 1-2 0V13c0-.58.253-1.047.539-1.38c.281-.33.63-.572.94-.742A1 1 0 0 0 12 9zm.999 4.011v-.004v.005zM12 15a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2H12z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}