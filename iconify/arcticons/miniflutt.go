package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Miniflutt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM22.03 18.956h6.683m-3.279 10.088V18.956"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.681 25.009a2.522 2.522 0 0 1 5.044 0v4.035m-5.044-6.557v6.557"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.725 25.009a2.522 2.522 0 1 1 5.043 0v4.035m11.868-10.088h6.683m-3.278 10.088V18.956"/>`),
		g.Group(children),
	)
}