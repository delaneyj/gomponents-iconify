package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretUpSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m7.5.265l7.207 7.761l-.733.68L7.5 1.736L1.026 8.707l-.733-.68L7.5.264Zm0 5.5l7.207 7.761l-.733.68L7.5 7.236l-6.474 6.972l-.733-.68L7.5 5.764Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}