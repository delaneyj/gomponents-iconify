package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.446 8.029L12 10.582l6.728-6.728a2 2 0 0 1 2.828 0l-12.11 12.11a4 4 0 1 1-1.414-1.414l2.554-2.553l-2.554-2.554a4 4 0 1 1 1.414-1.414Zm5.38 5.379l6.73 6.73a2 2 0 0 1-2.828 0l-5.317-5.316l1.415-1.414Zm-7.412 3.174a2 2 0 1 0-2.828 2.828a2 2 0 0 0 2.828-2.828Zm0-9.171a2 2 0 1 0-2.828-2.828A2 2 0 0 0 7.414 7.41Z"/>`),
		g.Group(children),
	)
}