package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Browser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 40H40a24.028 24.028 0 0 0-24 24v384a24.028 24.028 0 0 0 24 24h432a24.028 24.028 0 0 0 24-24V64a24.028 24.028 0 0 0-24-24ZM48 72h96v72H48Zm416 368H48V176h416Zm0-296H176V72h288Z"/>`),
		g.Group(children),
	)
}