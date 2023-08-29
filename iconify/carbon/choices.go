package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Choices(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 13V4h-9v2h5.586L16 14.586L7.414 6H13V4H4v9h2V7.414l9 9V26H4v2h24v-2H17v-9.586l9-9V13h2z"/>`),
		g.Group(children),
	)
}