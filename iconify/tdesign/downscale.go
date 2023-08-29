package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Downscale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h9v2H4v6H2V2Zm19.414 2l-5.027 5.027l3.527.02l-.011 2l-6.91-.04l-.04-6.91l2-.011l.02 3.526L20 2.585L21.414 4ZM2 12h4v2H4v2H2v-4Zm6 0h4v4h-2v-2H8v-2Zm14 1v9h-8v-2h6v-7h2ZM4 18v2h2v2H2v-4h2Zm8 0v4H8v-2h2v-2h2Z"/>`),
		g.Group(children),
	)
}