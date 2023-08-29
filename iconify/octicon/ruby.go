package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill-rule="evenodd" d="M13 6l-5 5V4h3l2 2zm3 0l-8 8l-8-8l4-4h8l4 4zm-8 6.5L14.5 6l-3-3h-7l-3 3L8 12.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}