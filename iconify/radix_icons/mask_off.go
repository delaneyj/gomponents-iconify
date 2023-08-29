package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaskOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 2h13v11H1V2ZM0 2a1 1 0 0 1 1-1h13a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1V2Zm4.875 5.5a2.625 2.625 0 1 1 5.25 0a2.625 2.625 0 0 1-5.25 0ZM7.5 3.875a3.625 3.625 0 1 0 0 7.25a3.625 3.625 0 0 0 0-7.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}