package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChecksThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m146.8 82.85l-89.6 88a4 4 0 0 1-5.6 0l-38.4-37.71a4 4 0 0 1 5.6-5.71l35.6 35l86.8-85.24a4 4 0 0 1 5.6 5.7Zm96-5.65a4 4 0 0 0-5.65 0l-86.8 85.24l-21.63-21.24a4 4 0 1 0-5.61 5.7l24.44 24a4 4 0 0 0 5.6 0l89.6-88a4 4 0 0 0 .1-5.7Z"/>`),
		g.Group(children),
	)
}