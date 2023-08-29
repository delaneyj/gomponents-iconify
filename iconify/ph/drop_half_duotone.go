package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropHalfDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 144a80 80 0 0 1-80 80V16s80 56 80 128Z" opacity=".2"/><path d="M174 47.75a254.19 254.19 0 0 0-41.45-38.3a8 8 0 0 0-9.18 0A254.19 254.19 0 0 0 82 47.75C54.51 79.32 40 112.6 40 144a88 88 0 0 0 176 0c0-31.4-14.51-64.68-42-96.25ZM56 144c0-50 42.26-92.71 64-111.4v182.94A72.08 72.08 0 0 1 56 144Zm80 71.54V32.6C157.74 51.29 200 94 200 144a72.08 72.08 0 0 1-64 71.54Z"/></g>`),
		g.Group(children),
	)
}