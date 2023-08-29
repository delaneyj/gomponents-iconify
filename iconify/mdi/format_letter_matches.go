package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatLetterMatches(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.14 4L6.43 16h1.93l.96-2.57h5.35l.97 2.57h1.93L12.86 4M12 6.29l2.03 5.42H9.96M20 14v4H4v-3H2v5h20v-6Z"/>`),
		g.Group(children),
	)
}