package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhonePageHeaderTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 4a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v1H5V4Zm0 2v10a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V6H5Z"/>`),
		g.Group(children),
	)
}