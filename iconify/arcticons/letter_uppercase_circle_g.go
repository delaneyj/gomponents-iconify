package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterUppercaseCircleG(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.128 19.996c0-4.644-3.878-8.38-8.565-8.134c-4.395.23-7.691 4.184-7.691 8.595v7.547c0 4.499 3.639 8.146 8.128 8.146h0c4.489 0 8.128-3.647 8.128-8.146H24"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}