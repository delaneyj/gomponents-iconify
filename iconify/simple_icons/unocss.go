package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unocss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.86 18.4a5.57 5.57 0 1 1 11.14 0a5.57 5.57 0 0 1-11.14 0Zm0-12.8A5.57 5.57 0 1 1 24 5.602v5.013a.557.557 0 0 1-.557.557H13.417a.557.557 0 0 1-.557-.557V5.6Zm-1.72 12.8A5.57 5.57 0 1 1 0 18.4v-5.014a.557.557 0 0 1 .557-.557h10.026a.557.557 0 0 1 .557.557V18.4Z"/>`),
		g.Group(children),
	)
}