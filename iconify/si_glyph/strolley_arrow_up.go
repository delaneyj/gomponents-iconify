package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StrolleyArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(1)"><ellipse cx="4.438" cy="14.459" rx="1.438" ry="1.459"/><ellipse cx="12.457" cy="14.419" rx="1.457" ry="1.419"/><path d="M15 10.016H3.953V3.778L1.781.41A.995.995 0 1 0 .174 1.584l1.849 2.901v5.531h-.055c-.553 0-1 .44-1 .984c0 .543.447.984 1 .984H15c.553 0 1-.441 1-.984a.991.991 0 0 0-1-.984z"/></g><path d="m13.768 3.757l-2.189-2.513a.811.811 0 0 0-1.143 0L8.247 3.757c-.314.315-.363.876-.049 1.19H10v1.922c0 .559.434 1.01.969 1.01c.535 0 .969-.451.969-1.01V4.947h1.781c.315-.315.364-.875.049-1.19z"/></g>`),
		g.Group(children),
	)
}