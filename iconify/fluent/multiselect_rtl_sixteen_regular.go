package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MultiselectRtlSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.854 2.854a.5.5 0 0 0-.708-.708L12.5 3.793l-.646-.647a.5.5 0 0 0-.708.708l1 1a.5.5 0 0 0 .708 0l2-2Zm0 9a.5.5 0 0 0-.708-.708L12.5 12.793l-.646-.647a.5.5 0 0 0-.708.708l1 1a.5.5 0 0 0 .708 0l2-2ZM1.5 3a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1h-8Zm0 4.5a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1h-8Zm0 4.5a.5.5 0 0 0 0 1h8a.5.5 0 0 0 0-1h-8Z"/>`),
		g.Group(children),
	)
}