package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardXSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 0h7v1h3v12.5a1.5 1.5 0 0 1-1.5 1.5h-10A1.5 1.5 0 0 1 1 13.5V1h3V0Zm1 1h5v1.5a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5V1Zm4.146 9.854L7.5 9.207l-1.646 1.646l-.708-.707L6.793 8.5L5.146 6.854l.708-.708L7.5 7.793l1.646-1.647l.708.708L8.207 8.5l1.647 1.646l-.708.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}