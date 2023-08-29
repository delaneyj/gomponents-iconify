package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27.188 4.875V5.97h-4.5V4.874H8.062V5.97h-4.5V4.874h-1v21.25h1V25.03h4.5v1.095h14.625V25.03h4.5v1.095h1.25V4.875h-1.25zM8.062 23.72h-4.5v-3.126h4.5v3.125zm0-4.44h-4.5v-3.124h4.5v3.125zm0-4.436h-4.5V11.72h4.5v3.124zm0-4.438h-4.5V7.28h4.5v3.126zm3.185 10.184V9.754l9.382 5.418l-9.383 5.418zm15.94 3.13h-4.5v-3.126h4.5v3.125zm0-4.44h-4.5v-3.124h4.5v3.125zm0-4.436h-4.5V11.72h4.5v3.124zm0-4.438h-4.5V7.28h4.5v3.126z"/>`),
		g.Group(children),
	)
}