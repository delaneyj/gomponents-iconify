package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CentMPassbook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.407 33.217H7.211a2.71 2.71 0 0 1-2.678-3.128L6.92 14.783H43.5l-2.533 16.242a2.59 2.59 0 0 1-2.56 2.192Zm-23.989-15.02h19.164m-14.996 3.521h10.828"/>`),
		g.Group(children),
	)
}