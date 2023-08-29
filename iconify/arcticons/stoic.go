package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stoic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1"/><rect width="5.571" height="7.381" x="21.992" y="22.39" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.785"/><circle cx="30.397" cy="18.978" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.397 22.39v7.382m-20.722-.623a3.133 3.133 0 0 0 2.291.623h.625a1.843 1.843 0 0 0 1.841-1.846h0a1.843 1.843 0 0 0-1.84-1.845h-1.25A1.843 1.843 0 0 1 9.5 24.236h0a1.843 1.843 0 0 1 1.841-1.846h.625a3.133 3.133 0 0 1 2.291.623m4.074-2.92v8.286a1.393 1.393 0 0 0 1.393 1.393h.418m-3.273-7.382h2.925M38.5 28.369a2.784 2.784 0 0 1-2.419 1.403h0a2.785 2.785 0 0 1-2.785-2.786v-1.81a2.785 2.785 0 0 1 2.785-2.786h0a2.784 2.784 0 0 1 2.416 1.398"/>`),
		g.Group(children),
	)
}