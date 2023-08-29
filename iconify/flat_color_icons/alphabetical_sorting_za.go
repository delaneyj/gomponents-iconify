package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphabeticalSortingZa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2196F3" d="M16.8 40h-5.3l-1.1 3H6.9l5.7-15.2h2.9L21.1 43h-3.2l-1.1-3zm-4.6-2.7H16l-1.9-5.7l-1.9 5.7zm.2-19.6H20v2.5H8.4v-1.9L16 7.5H8.4V5h11.4v1.7l-7.4 11z"/><path fill="#546E7A" d="M38 33V5h-4v28h-6l8 10l8-10z"/>`),
		g.Group(children),
	)
}