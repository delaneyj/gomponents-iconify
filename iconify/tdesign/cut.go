package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.414 4.586a2 2 0 1 0-2.828 2.828a2 2 0 0 0 2.828-2.828ZM3.172 3.172a4 4 0 0 1 6.274 4.86L12 10.586l8-8L21.414 4L9.446 15.968a4.002 4.002 0 0 1-6.274 4.86a4 4 0 0 1 4.86-6.274L10.586 12L8.032 9.446a4.002 4.002 0 0 1-4.86-6.274ZM15 13.586L21.414 20L20 21.414L13.586 15L15 13.586Zm-7.586 3a2 2 0 1 0-2.828 2.828a2 2 0 0 0 2.828-2.828Z"/>`),
		g.Group(children),
	)
}