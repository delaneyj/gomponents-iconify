package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M640 528q0 145-57 243.5T431 907l45 821q2 26-16 45t-44 19H224q-26 0-44-19t-16-45l45-821q-95-37-152-135.5T0 528q0-128 42.5-249.5T160 78.5T320 0t160 78.5t117.5 200T640 528z"/>`),
		g.Group(children),
	)
}