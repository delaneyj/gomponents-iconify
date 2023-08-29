package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Streamlabs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.316 14.25h18.672a6.512 6.512 0 0 1 6.512 6.512v13.78a6.512 6.512 0 0 1-6.512 6.511H15.71a3.907 3.907 0 0 1-3.908-3.907V20.762a6.512 6.512 0 0 1 6.512-6.512Zm7.218 12.051v7.011m8.91-7.011v7.011M19.107 6.947A14.607 14.607 0 0 0 4.5 21.553"/>`),
		g.Group(children),
	)
}