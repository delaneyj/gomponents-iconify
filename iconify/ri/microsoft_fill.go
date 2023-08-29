package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.501 3v8.5h-8.5V3h8.5Zm0 18h-8.5v-8.5h8.5V21Zm1-18h8.5v8.5h-8.5V3Zm8.5 9.5V21h-8.5v-8.5h8.5Z"/>`),
		g.Group(children),
	)
}