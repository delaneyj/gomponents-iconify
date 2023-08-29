package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M989.368 600.011L810.633 421.275L389.367 0L210.632 178.721l421.267 421.29l-421.267 421.267L389.355 1200l421.266-421.267L989.345 600.01h.023z"/>`),
		g.Group(children),
	)
}