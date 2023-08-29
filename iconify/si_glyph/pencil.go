package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.479 2.536L14.474.546c-.539-.537-1.383-.565-1.884-.064L3.47 9.616s-2.312 5.32-2.469 5.75c-.125.34.306.771.604.625c.48-.237 5.74-2.52 5.74-2.52l9.142-9.172c.502-.502.531-1.226-.008-1.763zM2.312 13.906l1.721-3.627l-.018.706l.998-.014l9.524-9.442l1.259 1.236l-9.184 8.965l-.227.752l.486.486l-3.811 1.656l-.748-.718z"/>`),
		g.Group(children),
	)
}