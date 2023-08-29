package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17 32C8.163 32 1 24.837 1 16S8.163 0 17 0s16 7.163 16 16s-7.163 16-16 16zm-1.286-4l-5.348-14.263l-2.16 9.84L15.714 28zM17.12 4.171l-5.863 7.132l5.863 14.983l5.897-14.983L17.12 4.17zM18.56 28l7.474-4.423l-2.125-9.84L18.56 28z"/>`),
		g.Group(children),
	)
}