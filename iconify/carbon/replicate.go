package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Replicate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 8h2V4a2.002 2.002 0 0 0-2-2h-4v2h4zM17 2h4v2h-4zm11 9h2v4h-2zm0 7v4h-4V10a2.002 2.002 0 0 0-2-2H10V4h4V2h-4a2.002 2.002 0 0 0-2 2v4H4a2.002 2.002 0 0 0-2 2v18a2.002 2.002 0 0 0 2 2h18a2.002 2.002 0 0 0 2-2v-4h4a2.002 2.002 0 0 0 2-2v-4zm-6 10H4V10h18z"/>`),
		g.Group(children),
	)
}