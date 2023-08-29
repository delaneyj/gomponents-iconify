package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AntennaBars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#3B88C3" d="M0 4c0-4 4-4 4-4h28s4 0 4 4v28s0 4-4 4H4s-4 0-4-4V4z"/><path fill="#FFF" d="M28 8h4v20h-4zm-6 4h4v16h-4zm-6 4h4v12h-4zm-6 4h4v8h-4zm-6 4h4v4H4z"/>`),
		g.Group(children),
	)
}