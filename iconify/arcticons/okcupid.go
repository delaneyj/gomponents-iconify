package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Okcupid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.5 19.6v20.9a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2h-33a2 2 0 0 0-2 2v12.1"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.901 29.4h0a3.396 3.396 0 0 1-3.438-4l.366-2.6a4.759 4.759 0 0 1 4.562-4h0a3.396 3.396 0 0 1 3.438 4l-.366 2.6a4.759 4.759 0 0 1-4.562 4Zm9.18-16l-2.249 16m.479-3.399l8.255-7.207m-5.627 4.913l4.896 5.669m9.821-1.991A4.716 4.716 0 0 1 33.9 29.4h0a3.396 3.396 0 0 1-3.438-4l.366-2.6a4.759 4.759 0 0 1 4.562-4h0a3.379 3.379 0 0 1 3.187 2.008"/>`),
		g.Group(children),
	)
}