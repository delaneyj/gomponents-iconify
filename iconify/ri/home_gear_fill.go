package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeGearFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 20a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-9H1l10.327-9.388a1 1 0 0 1 1.346 0L23 11h-3v9ZM8.592 13.808l-.991.572l1 1.733l.993-.573a3.5 3.5 0 0 0 1.405.811v1.145h2.002V16.35a3.5 3.5 0 0 0 1.405-.81l.992.572L16.4 14.38l-.991-.572a3.504 3.504 0 0 0 0-1.62l.991-.572l-1-1.734l-.993.574A3.499 3.499 0 0 0 13 9.644V8.5h-2.002v1.144a3.5 3.5 0 0 0-1.405.812l-.992-.574L7.6 11.616l.991.572a3.505 3.505 0 0 0 0 1.62Zm3.408.69a1.5 1.5 0 1 1-.002-3.001a1.5 1.5 0 0 1 .002 3Z"/>`),
		g.Group(children),
	)
}