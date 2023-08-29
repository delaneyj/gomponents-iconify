package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AntennaBars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M44 16h6v32h-6zm-10 8h6v24h-6zm-10 8h6v16h-6zm-10 8h6v8h-6z"/>`),
		g.Group(children),
	)
}