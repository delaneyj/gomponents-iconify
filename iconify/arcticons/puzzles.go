package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Puzzles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="38" height="38" x="5" y="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.18 9.2h6.09v6.09h-6.09zM20.96 20.95h6.09v6.09h-6.09zm11.22 11.97h6.09v6.09h-6.09zm0-20.67H12.77m0 23.71V12.25m19.41 23.71H12.77M24 20.95v-8.7m14.27 0H43m-7.77 17.18V18.57"/>`),
		g.Group(children),
	)
}