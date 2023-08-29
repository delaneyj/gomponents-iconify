package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="m12 14.676l-4.645 4.676A.5.5 0 0 1 6.5 19V5a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 .5.5v14a.5.5 0 0 1-.855.352L12 14.676Z" opacity=".2"/><path fill-rule="evenodd" d="M5.355 17.352L10 12.676l4.645 4.676A.5.5 0 0 0 15.5 17V3a.5.5 0 0 0-.5-.5H5a.5.5 0 0 0-.5.5v14a.5.5 0 0 0 .855.352Zm.145-1.565V3.5h9v12.287l-4.145-4.172a.5.5 0 0 0-.71 0L5.5 15.787Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}