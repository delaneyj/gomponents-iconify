package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeHeartFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 20a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-9H1l10.327-9.388a1 1 0 0 1 1.346 0L23 11h-3v9Zm-8-3l3.359-3.359a2.25 2.25 0 0 0-3.182-3.182l-.177.177l-.177-.177a2.25 2.25 0 1 0-3.182 3.182L12 17.001Z"/>`),
		g.Group(children),
	)
}