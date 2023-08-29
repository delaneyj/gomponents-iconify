package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xmos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 492.863L186.98 240L36.392 19.137H189.49l71.53 119.843l70.274-119.843h154.358L336.311 240L512 492.863H357.647L259.765 338.51L154.353 492.863z"/>`),
		g.Group(children),
	)
}