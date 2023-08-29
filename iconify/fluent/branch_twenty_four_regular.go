package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BranchTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 5.5a3.5 3.5 0 1 1 4.489 3.358a5.502 5.502 0 0 0 5.261 3.892h.33a3.501 3.501 0 0 1 6.92.75a3.5 3.5 0 0 1-6.92.75h-.33a6.987 6.987 0 0 1-5.5-2.67v3.5A3.501 3.501 0 0 1 7.5 22a3.5 3.5 0 0 1-.75-6.92V8.92A3.501 3.501 0 0 1 4 5.5Zm3.5-2a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm0 13a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm8-3a2 2 0 1 0 4 0a2 2 0 0 0-4 0Z"/>`),
		g.Group(children),
	)
}