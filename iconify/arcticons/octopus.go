package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Octopus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.974 19.715c1.323-6.796 10.497-12.096 13.937-7.013c2.646 3.909-3.908 9.201-14.071 8.9c-10.184-.301-15.548-1.4-19.845 1.383c-4.27 2.767-5.773 7.938-3.307 11.486s10.583 4.39 16.176-1.022S28.084 24 28.084 24"/>`),
		g.Group(children),
	)
}