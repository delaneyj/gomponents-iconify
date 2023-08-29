package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.07 32.207h9.353M14.069 13.5h9.354m-9.354 9.354h6.08M14.07 13.5v18.707m19.86-4.66a4.675 4.675 0 0 1-4.66 4.66h0a4.675 4.675 0 0 1-4.661-4.66v-3.03a4.675 4.675 0 0 1 4.661-4.661h0a4.675 4.675 0 0 1 4.661 4.66m0-4.66V38.5"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}