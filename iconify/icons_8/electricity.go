package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Electricity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m18.875 4l-6.906 7.313l-.845.906l1.063.624l7.187 4.344L12 24.563V20h-2v8h8v-2h-4.563l8.282-8.28l.905-.907l-1.094-.657l-7.25-4.375l6.033-6.405L18.875 4z"/>`),
		g.Group(children),
	)
}