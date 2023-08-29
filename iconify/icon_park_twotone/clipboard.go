package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTClipboard0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M17 7h-7a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-7"/><path fill="#555" d="M17 4h14v6H17z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTClipboard0)"/>`),
		g.Group(children),
	)
}