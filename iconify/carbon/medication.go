package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medication(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 2H8a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2v18a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V10a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2ZM10 14h3v10h-3Zm12 14H10v-2h5V12h-5v-2h12ZM8 8V4h16v4Z"/>`),
		g.Group(children),
	)
}