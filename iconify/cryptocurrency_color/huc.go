package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Huc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#ffc018"/><path fill="#fff" d="M11.5 14.5h9V10h3v16h-3v-8.5h-9V22h-3V6h3z"/></g>`),
		g.Group(children),
	)
}