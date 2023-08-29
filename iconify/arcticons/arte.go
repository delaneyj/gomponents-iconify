package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arte(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.48 5.5a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33.04a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.468 24.326a3.72 3.72 0 0 1 3.72-3.721h0m-3.72 0v9.859m18.544-1.877a3.72 3.72 0 0 1-3.232 1.877h0a3.72 3.72 0 0 1-3.72-3.72v-2.418a3.72 3.72 0 0 1 3.72-3.721h0a3.72 3.72 0 0 1 3.72 3.72v1.21h-7.441m-3.882-7.999v11.068a1.86 1.86 0 0 0 1.86 1.86h.558m-4.371-9.859h3.906m-12.189 6.139a3.72 3.72 0 0 1-3.72 3.72h0a3.72 3.72 0 0 1-3.721-3.72v-2.418a3.72 3.72 0 0 1 3.72-3.721h0a3.72 3.72 0 0 1 3.72 3.72m.001 6.139v-9.859"/>`),
		g.Group(children),
	)
}