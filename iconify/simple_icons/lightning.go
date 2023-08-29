package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lightning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0L1.75 6v12L12 24l10.25-6V6zm-1.775 18l1.08-4.657l-2.428-2.397L13.79 6l-1.082 4.665l2.414 2.384z"/>`),
		g.Group(children),
	)
}