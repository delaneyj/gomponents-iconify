package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Save(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m472.971 122.344l-99.315-99.315A23.838 23.838 0 0 0 356.687 16H56a24.028 24.028 0 0 0-24 24v432a24.028 24.028 0 0 0 24 24h400a24.028 24.028 0 0 0 24-24V139.313a23.838 23.838 0 0 0-7.029-16.969ZM320 48v96H176V48Zm128 416H64V48h80v128h208V48h1.373L448 142.627Z"/><path fill="currentColor" d="M252 224a92 92 0 1 0 92 92a92.1 92.1 0 0 0-92-92Zm0 152a60 60 0 1 1 60-60a60.068 60.068 0 0 1-60 60Z"/>`),
		g.Group(children),
	)
}