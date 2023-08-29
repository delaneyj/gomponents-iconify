package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DebugStepOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 1h-.542L3.553 4.905l1.061 1.06l2.637-2.61v6.177h1.498V3.355l2.637 2.61l1.061-1.06L8.542 1H8zm1.956 12.013a2 2 0 1 1-4 0a2 2 0 0 1 4 0z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}