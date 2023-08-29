package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Datadex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.838 26.332a6.688 6.688 0 0 0-6.668-6.668h0a6.688 6.688 0 0 0-6.669 6.668v4.335a6.688 6.688 0 0 0 6.669 6.669h0a6.688 6.688 0 0 0 6.669-6.67m0 6.67V10.66m7.99 26.678V10.662h4.334A13.377 13.377 0 0 1 43.501 24h0a13.377 13.377 0 0 1-13.338 13.338Z"/><ellipse cx="34.748" cy="23.873" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.301" ry="3.165"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.145 24h5.346m12.006 0h-5.342"/>`),
		g.Group(children),
	)
}