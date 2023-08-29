package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Passports(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M18.5 5.5h2v18h-17v-20c8.5-1 14.75-3 14.75-3h.25v5Zm0 0H6m10.5 9A4.5 4.5 0 0 0 12 10m4.5 4.5A4.5 4.5 0 0 1 12 19m4.5-4.5c-.5.5-2 1-4.5 1s-4-.5-4.5-1M12 10a4.5 4.5 0 0 0-4.5 4.5M12 10c.966 0 1.75 2.015 1.75 4.5S12.966 19 12 19m0-9c-.966 0-1.75 2.015-1.75 4.5S11.034 19 12 19m-4.5-4.5A4.5 4.5 0 0 0 12 19"/>`),
		g.Group(children),
	)
}