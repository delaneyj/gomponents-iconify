package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FemaleSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 15.5a7.5 7.5 0 1 1 0-15a7.5 7.5 0 0 1 0 15Zm0 0V24m-4-4h8"/>`),
		g.Group(children),
	)
}