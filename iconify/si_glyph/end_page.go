package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EndPage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m8.327 11.886l4.447-4.94a.65.65 0 0 0-.002-.849l-2.841-.005V1.068c0-.553-.437-1-.976-1H7.004a.987.987 0 0 0-.976 1v5.02l-2.95-.005a.652.652 0 0 0 .004.848l4.485 4.954a.501.501 0 0 0 .76.001zm5.591 2.948c0 .552-.437 1-.973 1H3.056c-.537 0-.973-.448-.973-1V14c0-.552.436-1 .973-1h9.889c.536 0 .973.448.973 1v.834z"/>`),
		g.Group(children),
	)
}