package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoinInner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 6a9.935 9.935 0 0 0-4 .842a9.999 9.999 0 1 0 0 18.318A9.998 9.998 0 1 0 20 6Zm-8 18a8 8 0 1 1 1.757-15.798a9.973 9.973 0 0 0 0 15.598A7.992 7.992 0 0 1 12 24Zm8 0a7.992 7.992 0 0 1-1.757-.2a9.973 9.973 0 0 0 0-15.599A7.997 7.997 0 1 1 20 24Z"/>`),
		g.Group(children),
	)
}