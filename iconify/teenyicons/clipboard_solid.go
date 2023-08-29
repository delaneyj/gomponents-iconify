package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 0H4v1H1v12.5A1.5 1.5 0 0 0 2.5 15h10a1.5 1.5 0 0 0 1.5-1.5V1h-3V0Zm-1 1H5v1.5a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}