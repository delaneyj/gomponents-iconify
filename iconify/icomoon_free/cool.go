package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cool(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16zM8 1.5a6.5 6.5 0 1 1 0 13a6.5 6.5 0 0 1 0-13zM12.5 4c.275 0 .5.225.5.5V6c0 .55-.45 1-1 1h-2c-.55 0-1-.45-1-1H7c0 .55-.45 1-1 1H4c-.55 0-1-.45-1-1V4.5c0-.275.225-.5.5-.5h3c.275 0 .5.225.5.5V5h2v-.5c0-.275.225-.5.5-.5h3zM8 12a3.996 3.996 0 0 0 3.43-1.942l.857.515a4.996 4.996 0 0 1-6.406 1.957l.518-.864c.49.214 1.031.334 1.6.334z"/>`),
		g.Group(children),
	)
}