package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeGearLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 21H5a1 1 0 0 1-1-1v-9H1l10.327-9.388a1 1 0 0 1 1.346 0L23 11h-3v9a1 1 0 0 1-1 1ZM6 19h12V9.158l-6-5.455l-6 5.455V19Zm2.591-5.191a3.508 3.508 0 0 1 0-1.622l-.991-.572l1-1.732l.991.573a3.494 3.494 0 0 1 1.404-.812V8.5h2v1.144c.532.159 1.01.44 1.404.812l.991-.573l1 1.732l-.991.572a3.508 3.508 0 0 1 0 1.622l.991.572l-1 1.731l-.991-.572a3.494 3.494 0 0 1-1.404.811v1.145h-2V16.35a3.495 3.495 0 0 1-1.404-.811l-.991.573l-1-1.732l.991-.572Zm3.404.688a1.5 1.5 0 1 0 0-2.998a1.5 1.5 0 0 0 0 2.998Z"/>`),
		g.Group(children),
	)
}