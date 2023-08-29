package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FenceOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12H4v4h12m4 0v-4h-4M6 16v4h4v-4m0-4v-2m0-4L8 4M6 6v6m8 4v4h4v-2m0-6V6l-2-2l-2 2v4M3 3l18 18"/>`),
		g.Group(children),
	)
}