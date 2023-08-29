package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyEe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.07 24a11.007 11.007 0 0 0 3.963-8.467C35.033 9.44 30.093 4.5 24 4.5S12.967 9.44 12.967 15.533c0 3.403 1.542 6.443 3.964 8.467a11.007 11.007 0 0 0-3.964 8.467C12.967 38.56 17.907 43.5 24 43.5s11.033-4.94 11.033-11.033A11.01 11.01 0 0 0 31.069 24Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.66 32.467h3.051m1.629 4.68h-4.68v-9.36h4.68m-4.68-12.254h3.051m1.629 4.68h-4.68v-9.36h4.68"/>`),
		g.Group(children),
	)
}