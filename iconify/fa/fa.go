package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1504 512v839q0 48-49 62q-174 52-338 52q-73 0-215.5-29.5T674 1406q-164 0-370 48v338H144V424q-63-25-101-81T5 219q0-91 64-155T224 0t155 64t64 155q0 68-38 124t-101 81v68q190-44 343-44q99 0 198 15q14 2 111.5 22.5T1106 506q77 0 165-18q11-2 80-21t89-19q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}