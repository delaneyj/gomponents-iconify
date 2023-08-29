package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileSettingsLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.595 12.812a3.51 3.51 0 0 1 0-1.623l-.992-.573l1-1.732l.992.573A3.496 3.496 0 0 1 11 8.645V7.5h2v1.145c.532.158 1.012.44 1.405.812l.992-.573l1 1.732l-.992.572a3.507 3.507 0 0 1 0 1.623l.992.573l-1 1.732l-.992-.573a3.495 3.495 0 0 1-1.405.812V16.5h-2v-1.145a3.495 3.495 0 0 1-1.405-.812l-.992.573l-1-1.732l.992-.572ZM12 13.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3ZM15 4H5v16h14V8h-4V4ZM3 2.992C3 2.444 3.447 2 3.998 2H16l5 5v13.992A1 1 0 0 1 20.007 22H3.993A1 1 0 0 1 3 21.008V2.992Z"/>`),
		g.Group(children),
	)
}