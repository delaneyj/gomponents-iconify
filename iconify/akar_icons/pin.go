package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 10.726l-3 .505L11.23 2l-.504 3M12 16.881l-.77 2.042l7.693-7.692l-2.042.769m-1.804 3.077L21 21M3.538 9.692l6.154-6.154l.236.341a52.22 52.22 0 0 0 7.376 8.518l.235.218l-4.924 4.923l-.218-.234A52.22 52.22 0 0 0 3.88 9.928l-.34-.236Z"/>`),
		g.Group(children),
	)
}