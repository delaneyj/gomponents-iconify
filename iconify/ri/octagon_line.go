package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OctagonLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.936 2.501l5.565 5.565v7.87l-5.565 5.565h-7.87l-5.565-5.565v-7.87l5.565-5.565h7.87Zm-.828 2H8.894L4.501 8.894v6.214L8.894 19.5h6.214l4.393-4.393V8.894l-4.393-4.393Z"/>`),
		g.Group(children),
	)
}