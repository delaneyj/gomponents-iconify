package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rax(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#F76927" d="m.61.667l62.848 64.059l103.504-.764s23.118 1.788 23.118 30.211c0 28.424-19.61 30.47-19.61 30.47l-42.732 1.146l58.56 61.224S256 171.199 256 94.413S176.607 0 176.607 0L.611.667ZM0 126.313h89.65l127.598 128.36l-90.652 1.76L0 126.313Z"/>`),
		g.Group(children),
	)
}