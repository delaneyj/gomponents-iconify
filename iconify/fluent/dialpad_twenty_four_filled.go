package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialpadTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 17.75a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm0-5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm4.996 0a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm-9.992 0a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm4.996-5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm4.996 0a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm-9.992 0a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm4.996-5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm4.996 0a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm-9.992 0a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Z"/>`),
		g.Group(children),
	)
}