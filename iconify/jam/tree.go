package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.728 5.547l-.632-.446l-.167-.755a3.002 3.002 0 0 0-5.858 0l-.167.755l-.632.446A3 3 0 1 0 6.2 10.75l.801-.35l.8.35a3 3 0 0 0 2.927-5.204zM8 12.9V19a1 1 0 0 1-2 0v-6.1a5 5 0 0 1-3.882-8.987a5.002 5.002 0 0 1 9.764 0A5 5 0 0 1 8 12.9z"/>`),
		g.Group(children),
	)
}