package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightHouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.042 0S4.073 1.93 4.073 3h7.938c0-1.07-3.969-3-3.969-3zM12 5v.982L16 7V4l-4 1zM3.958 5L0 4v3l3.958-1V5zM5 9.969h6V9h1V8H4v1h1v.969zM5 14h6v2H5zm0-3h6v2H5zm0-7h.968v3H5zm5 0h.968v3H10zM7 4h1.984v3H7z"/>`),
		g.Group(children),
	)
}