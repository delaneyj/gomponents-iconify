package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaceholderDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 48v160a8 8 0 0 1-8 8H48a8 8 0 0 1-8-8V48a8 8 0 0 1 8-8h160a8 8 0 0 1 8 8Z" opacity=".2"/><path d="M224 48a16 16 0 0 0-16-16H48a15.91 15.91 0 0 0-10.66 4.1a9.08 9.08 0 0 0-1.24 1.24A15.91 15.91 0 0 0 32 48v160a16 16 0 0 0 16 16h160a15.91 15.91 0 0 0 10.66-4.1a7.35 7.35 0 0 0 .65-.59a6 6 0 0 0 .58-.65A15.87 15.87 0 0 0 224 208Zm-16 148.7L59.31 48H208ZM48 59.31L196.69 208H48Z"/></g>`),
		g.Group(children),
	)
}