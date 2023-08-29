package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M63 1408H0V0h63v1408zm63-1H94V0h32v1407zm94 0h-31V0h31v1407zm157 0h-31V0h31v1407zm157 0h-62V0h62v1407zm126 0h-31V0h31v1407zm63 0h-31V0h31v1407zm63 0h-31V0h31v1407zm157 0h-63V0h63v1407zm157 0h-63V0h63v1407zm126 0h-63V0h63v1407zm126 0h-63V0h63v1407zm94 0h-63V0h63v1407zm189 0h-94V0h94v1407zm63 0h-32V0h32v1407zm94 1h-63V0h63v1408z"/>`),
		g.Group(children),
	)
}