package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WidgetThreeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.5 1.75a4.75 4.75 0 1 0 0 9.5a4.75 4.75 0 0 0 0-9.5ZM3.25 6.5a3.25 3.25 0 1 1 6.5 0a3.25 3.25 0 0 1-6.5 0Zm14.25 6.25a4.75 4.75 0 1 0 0 9.5a4.75 4.75 0 0 0 0-9.5Zm-3.25 4.75a3.25 3.25 0 1 1 6.5 0a3.25 3.25 0 0 1-6.5 0Zm-1.5-11a4.75 4.75 0 1 1 9.5 0a4.75 4.75 0 0 1-9.5 0Zm4.75-3.25a3.25 3.25 0 1 0 0 6.5a3.25 3.25 0 0 0 0-6.5Zm-11 9.5a4.75 4.75 0 1 0 0 9.5a4.75 4.75 0 0 0 0-9.5ZM3.25 17.5a3.25 3.25 0 1 1 6.5 0a3.25 3.25 0 0 1-6.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}