package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpamThreeLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.935 2.501L21.5 8.066v7.87l-5.565 5.565h-7.87L2.5 15.936v-7.87l5.565-5.565h7.87Zm-.828 2H8.893L4.5 8.894v6.214L8.893 19.5h6.214l4.393-4.393V8.894l-4.393-4.393Zm-7.108 6.5h8v2H8v-2Z"/>`),
		g.Group(children),
	)
}