package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaStop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M408 80H104a24.027 24.027 0 0 0-24 24v304a24.027 24.027 0 0 0 24 24h304a24.027 24.027 0 0 0 24-24V104a24.027 24.027 0 0 0-24-24Zm-8 320H112V112h288Z"/>`),
		g.Group(children),
	)
}