package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Freedcamviewer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 17.435V37.5a2 2 0 0 0 2 2h35a2 2 0 0 0 2-2v-27a2 2 0 0 0-2-2H13.435"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.777 35.307l6.275-12.736l5.785 9.047l9.045-15.581l9.341 19.27H8.777zm12.06-3.689l-2.135 3.689M8.5 5.52l6.978 6.98L8.5 19.477l-6.978-6.979z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.435 7.565v9.869H3.566V7.565z"/>`),
		g.Group(children),
	)
}