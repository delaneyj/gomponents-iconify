package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Looks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M234.5 149q61.5 0 105.5 44t44 106h-43q0-44-31-75.5T235 192t-75.5 31.5T128 299H85q0-62 44-106t105.5-44zm.5-85q97 0 165.5 69T469 299h-42q0-80-56.5-136t-136-56T99 163T43 299H0q0-97 69-166t166-69z"/>`),
		g.Group(children),
	)
}