package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShowerDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="m184 72l-22.09 129.31a8 8 0 0 1-13.54 4.34l-98-98a8 8 0 0 1 4.34-13.54Z" opacity=".2"/><path d="M64 236a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm20-44a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm-64 0a12 12 0 1 0 12 12a12 12 0 0 0-12-12Zm32-32a12 12 0 1 0 12 12a12 12 0 0 0-12-12ZM256 40a8 8 0 0 1-8 8h-28.69l-27.85 27.86l-21.66 126.79a16 16 0 0 1-27.09 8.66l-98-98a16 16 0 0 1 8.69-27.1l126.74-21.67L208 36.69A15.86 15.86 0 0 1 219.31 32H248a8 8 0 0 1 8 8Zm-81.79 41.79L56 102l98 98Z"/></g>`),
		g.Group(children),
	)
}