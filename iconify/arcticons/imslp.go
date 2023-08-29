package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Imslp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.5 20.108v7.784m1.963-.097v-7.687l3.892 7.784l3.892-7.784v7.784m1.79-.876c.487.584 1.07.876 1.947.876h1.167c1.07 0 1.946-.876 1.946-1.946h0c0-1.07-.875-1.946-1.946-1.946h-1.265a1.952 1.952 0 0 1-1.946-1.946h0c0-1.07.876-1.946 1.946-1.946h1.168c.876 0 1.46.194 1.946.875m1.68-.875v7.784h3.892m1.771 0v-7.784h2.53c1.46 0 2.627 1.168 2.627 2.627s-1.167 2.627-2.627 2.627h-2.53"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}