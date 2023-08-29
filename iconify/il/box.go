package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Box(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M729 241q12 16 11 33l-20 261q-2 22-24 22H44q-9 0-15-6t-8-16L1 274q-1-17 10-33L198 14q10-13 27-13h291q16 0 27 13zm-85 73q0-5-3-9t-9-3H109q-5 0-8 3t-4 9l11 140q0 10 11 10h502q10 0 12-10z"/>`),
		g.Group(children),
	)
}