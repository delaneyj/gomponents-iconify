package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SaveAs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.04 1.33L12.71 3l.29.71v.33h-.5l-.5.5v-.83l-1.67-1.67H10v4H4v-4H2v10h3l-.5 1H2l-1-1v-10l1-1h8.33l.71.29zM7 5h2V2H7v3zm6.5 0L15 6.5l-.02.69l-5.5 5.5l-.13.12l-.37.37l-.1.09l-3 1.5l-.67-.67l1.5-3l.09-.1l.37-.37l.12-.13l5.5-5.5h.71zm-6.22 7.24l-.52 1l1.04-.48l-.52-.52zm.69-1.03l.79.79l5.15-5.15l-.79-.79l-5.15 5.15z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}