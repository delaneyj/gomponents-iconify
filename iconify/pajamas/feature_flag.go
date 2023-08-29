package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FeatureFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 5v6a3.5 3.5 0 1 1-7 0V5a3.5 3.5 0 1 1 7 0ZM3 5a5 5 0 0 1 10 0v6a5 5 0 0 1-10 0V5Zm5 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}