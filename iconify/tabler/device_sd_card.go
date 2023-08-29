package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeviceSdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21h10a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-6.172a2 2 0 0 0-1.414.586L5.586 7.414A2 2 0 0 0 5 8.828V19a2 2 0 0 0 2 2zm6-15v2m3-2v2m-6-1v1"/>`),
		g.Group(children),
	)
}