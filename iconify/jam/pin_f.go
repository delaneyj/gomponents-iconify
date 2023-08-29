package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PinF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.374 12.268l-5.656 5.657A1 1 0 1 1 .303 16.51l5.657-5.656L1.718 6.61A6.992 6.992 0 0 1 7.9 4.67L11.617.954a2 2 0 0 1 2.828 0l2.829 2.828a2 2 0 0 1 0 2.829l-3.716 3.716a6.992 6.992 0 0 1-1.941 6.183l-4.243-4.242z"/>`),
		g.Group(children),
	)
}