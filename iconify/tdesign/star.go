package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .63l2.903 8.35l8.839.181l-7.045 5.341l2.56 8.462L12 17.914l-7.256 5.05l2.56-8.462L.259 9.161l8.839-.18L12 .63Zm0 6.092l-1.47 4.23l-4.478.091l3.569 2.706l-1.297 4.287L12 15.478l3.676 2.558l-1.296-4.287l3.568-2.706l-4.477-.09L12 6.721Z"/>`),
		g.Group(children),
	)
}