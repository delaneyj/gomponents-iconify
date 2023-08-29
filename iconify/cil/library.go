package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Library(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M472 16H168a24 24 0 0 0-24 24v304a24 24 0 0 0 24 24h304a24 24 0 0 0 24-24V40a24 24 0 0 0-24-24Zm-8 320H176V48h288Z"/><path fill="currentColor" d="M112 400V80H80v328a24 24 0 0 0 24 24h328v-32Z"/><path fill="currentColor" d="M48 464V144H16v328a24 24 0 0 0 24 24h328v-32Z"/>`),
		g.Group(children),
	)
}