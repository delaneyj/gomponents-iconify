package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Previous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M274.3 262.5L512 381.4V143.6L274.3 262.5zm-237.7 0l237.7 118.9V143.6L36.6 262.5zM0 143.6v237.7h36.6V143.6H0z"/>`),
		g.Group(children),
	)
}