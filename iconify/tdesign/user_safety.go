package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserSafety(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Zm10.5 8v3.634a1 1 0 0 0 .453.837L19 21.308l2.047-1.337a1 1 0 0 0 .453-.837V15.5h-5Zm-2-2h9v5.634a3 3 0 0 1-1.36 2.511L19 23.697l-3.14-2.052a3 3 0 0 1-1.36-2.511V13.5ZM8 16a4 4 0 0 0-4 4h8.55v2H2v-2a6 6 0 0 1 6-6h4.5v2H8Z"/>`),
		g.Group(children),
	)
}