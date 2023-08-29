package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myjio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.66 16.34v11.49a3.82 3.82 0 0 1-3.83 3.83h0A3.83 3.83 0 0 1 12 27.83v-1.27"/><rect width="7.66" height="10.15" x="28.34" y="21.51" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.83"/><circle cx="24.1" cy="18.07" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.1 21.51v10.15"/>`),
		g.Group(children),
	)
}