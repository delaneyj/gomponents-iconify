package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControlPlatform(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 .856l10 5.556v11.176l-10 5.556l-10-5.556V6.412L12 .856Zm-8 8.04v7.515l7 3.89v-7.699L4 8.896ZM13 20.3l7-3.889V8.896l-7 3.706V20.3Zm-1-9.432L19.12 7.1L12 3.144L4.88 7.099L12 10.87Z"/>`),
		g.Group(children),
	)
}