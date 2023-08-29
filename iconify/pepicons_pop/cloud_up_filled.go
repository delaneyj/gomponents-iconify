package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudUpFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 14v-3.322l.752.587a1 1 0 0 0 1.231-1.576l-2.304-1.8a1 1 0 0 0-.673-.259a1 1 0 0 0-.674.259l-2.304 1.8a1 1 0 0 0 1.231 1.576l.741-.578V14H6a4 4 0 0 1 0-8h.126C6.57 4.275 8.136 3 10 3h1c1.9 0 3.49 1.325 3.899 3.101A4.002 4.002 0 0 1 14 14h-2.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}