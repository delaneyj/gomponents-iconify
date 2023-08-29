package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AButtonBloodType(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M15.907 7.689a1.75 1.75 0 0 1 1.63 1.09L22.886 21.9a1.75 1.75 0 1 1-3.241 1.322l-.734-1.801a.125.125 0 0 0-.116-.078h-5.613a.125.125 0 0 0-.116.079l-.704 1.782a1.75 1.75 0 0 1-3.255-1.286l5.182-13.122a1.75 1.75 0 0 1 1.618-1.107Zm1.475 9.982l-1.322-3.244a.125.125 0 0 0-.232 0l-1.281 3.245a.125.125 0 0 0 .116.17h2.603c.09 0 .15-.09.116-.171Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}