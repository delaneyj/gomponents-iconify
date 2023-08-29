package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextUnderlineSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.5 2a.5.5 0 0 1 .5.5V8c0 1.624 1.376 3 3 3s3-1.376 3-3V2.5a.5.5 0 0 1 1 0V8c0 2.176-1.824 4-4 4s-4-1.824-4-4V2.5a.5.5 0 0 1 .5-.5ZM4 13.5a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Z"/>`),
		g.Group(children),
	)
}