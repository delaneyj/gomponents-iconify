package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.225 5.767V3.086L0 9.542l7.225 6.691v-2.777L3 9.542l4.225-3.775zm5 1.186V3.086L5 9.542l7.225 6.691v-4.357c3.292 0 5.291.422 7.775 4.81c0-.001-.368-9.733-7.775-9.733z"/>`),
		g.Group(children),
	)
}