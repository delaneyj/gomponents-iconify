package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YleAreena(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.721 29.113l-3.929-9.592m7.484 0l-4.678 13.031A2.184 2.184 0 0 1 15.54 34h-.561m9.639-18v12.282m8.112-1.772a3.094 3.094 0 0 1-2.713 1.58h0a3.184 3.184 0 0 1-3.192-3.158v-2.054a3.184 3.184 0 0 1 3.192-3.158h0a3.184 3.184 0 0 1 3.191 3.158v1.106h-6.383"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}