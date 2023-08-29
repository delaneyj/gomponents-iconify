package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oeffiplans(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5A21.5 21.5 0 1 0 45.5 24A21.51 21.51 0 0 0 24 2.5ZM6.93 21.27h12.64m7.59 0h13.91M6.93 26.43h12.64m7.59 0h13.91m-17.7-11.46V7.38"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.37 32.67v2.93l-5.03 5.02"/><rect width="7.59" height="17.71" x="19.57" y="14.97" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}