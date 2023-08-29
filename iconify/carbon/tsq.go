package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tsq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M19 26h11v2H19zm0-4h11v2H19zm11-2H19v-8h11v8zm-9-2h7v-4h-7v4zM19 8h11v2H19zm0-4h11v2H19zm-8.707 14.707L8 16.414V12h2v3.586l1.707 1.707l-1.414 1.414z"/><path fill="currentColor" d="M9 24c-4.411 0-8-3.589-8-8s3.589-8 8-8s8 3.589 8 8s-3.589 8-8 8Zm0-14c-3.308 0-6 2.692-6 6s2.692 6 6 6s6-2.692 6-6s-2.692-6-6-6Z"/>`),
		g.Group(children),
	)
}