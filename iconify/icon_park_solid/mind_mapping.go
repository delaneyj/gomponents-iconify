package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MindMapping(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMindMapping0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M8 28a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M42 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 18a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0 18a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/><path stroke-linecap="round" d="M32 6H20v36h12M12 24h20"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMindMapping0)"/>`),
		g.Group(children),
	)
}