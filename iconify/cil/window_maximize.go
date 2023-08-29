package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WindowMaximize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 48H40.335a24.027 24.027 0 0 0-24 24v384a24.027 24.027 0 0 0 24 24H472a24.027 24.027 0 0 0 24-24V72a24.027 24.027 0 0 0-24-24Zm-8 32v71.981l-415.665-.491V80ZM48.335 448V183.49l415.665.491V448Z"/>`),
		g.Group(children),
	)
}