package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeHeartLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 20a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-9H1l10.327-9.388a1 1 0 0 1 1.346 0L23 11h-3v9Zm-2-1V9.158l-6-5.455l-6 5.455V19h12Zm-6-2l-3.359-3.359a2.25 2.25 0 1 1 3.182-3.182l.177.177l.177-.177a2.25 2.25 0 1 1 3.182 3.182L12 17.001Z"/>`),
		g.Group(children),
	)
}