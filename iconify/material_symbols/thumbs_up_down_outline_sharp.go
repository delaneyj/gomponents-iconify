package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsUpDownOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 14V5.4L5.4 0l1.225 1.225L5.8 5H12v2.575L9.275 14H0Zm2-2h5.95L10 7.15V7H3.35l.6-2.7L2 6.2V12Zm16.6 12l-1.225-1.225L18.2 19H12v-2.575L14.725 10H24v8.6L18.6 24Zm1.45-4.3L22 17.8V12h-5.95L14 16.85V17h6.65l-.6 2.7ZM2 12V6.2V12Zm20 5.8V12v5.8Z"/>`),
		g.Group(children),
	)
}