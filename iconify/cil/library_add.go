package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LibraryAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M48 464V144H16v328a24.027 24.027 0 0 0 24 24h328v-32H48Z"/><path fill="currentColor" d="M144 400h-32V80H80v328a24.027 24.027 0 0 0 24 24h328v-32H144Z"/><path fill="currentColor" d="M472 16H168a24.027 24.027 0 0 0-24 24v304a24.027 24.027 0 0 0 24 24h304a24.027 24.027 0 0 0 24-24V40a24.027 24.027 0 0 0-24-24Zm-8 320H176V48h288Z"/><path fill="currentColor" d="M304 288h32v-84h84v-32h-84V88h-32v84h-84v32h84v84z"/>`),
		g.Group(children),
	)
}