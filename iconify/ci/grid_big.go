package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridBig(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 19h-6v-6h6v6Zm-8 0H5v-6h6v6Zm8-8h-6V5h6v6Zm-8 0H5V5h6v6Z"/>`),
		g.Group(children),
	)
}