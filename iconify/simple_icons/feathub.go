package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Feathub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.571 0v6.857h6.858V0zM0 8.571v6.858h24V8.57zm8.571 8.572V24h6.858v-6.857z"/>`),
		g.Group(children),
	)
}