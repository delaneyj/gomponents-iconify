package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalCellularFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M384.78 88h-.1v.1L208.78 264h-.1v.1L16 456.78V496h480V16h-39.22Zm-264.1 376H54.034l66.647-66.646Zm88 0h-56v-98.646l56-56Zm88 0h-56V277.354l56-56Zm88 0h-56V189.354l56-56Zm79.32 0h-47.319V101.354L464 54.034Z"/>`),
		g.Group(children),
	)
}