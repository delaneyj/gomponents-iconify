package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoSemicircles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTwoSemicircles0"><path fill="#fff" fill-rule="evenodd" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M44 25c0-11.046-8.954-20-20-20S4 13.954 4 25h40Zm-30 7c0 5.523 4.477 10 10 10s10-4.477 10-10H14Z" clip-rule="evenodd"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTwoSemicircles0)"/>`),
		g.Group(children),
	)
}