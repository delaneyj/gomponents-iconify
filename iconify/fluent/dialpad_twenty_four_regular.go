package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DialpadTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 17.75a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5Zm0-5a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5Zm4.996 0a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5Zm-9.992 0a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5Zm4.996-5a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5Zm4.996 0a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5Zm-9.992 0a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5Zm4.996-5a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5Zm4.996 0a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5Zm-9.992 0a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5Z"/>`),
		g.Group(children),
	)
}