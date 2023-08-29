package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picasa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18.983 14.377c-1.533 3.166-4.617 5.3-8.07 5.623H9.075a9.967 9.967 0 0 1-3.174-.843v-4.78ZM4.853 9.964v8.615A10.045 10.045 0 0 1 .72 13.726l1.889-1.72l.378-.344l1.864-1.698Zm10.324-8.491C18.152 3.287 20 6.539 20 10.02c0 1.122-.195 2.232-.567 3.3h-4.256Zm-9.734-.361l1.59 1.446l.563.512l1.63 1.484L4.71 8.668l-.488.444l-3.859 3.512A10.036 10.036 0 0 1 0 10.021a9.973 9.973 0 0 1 5.443-8.91ZM9.999 0c1.433 0 2.82.304 4.124.898V7.58L6.507.653A9.91 9.91 0 0 1 10 0Z"/>`),
		g.Group(children),
	)
}