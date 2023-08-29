package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InterfaceEditAttachmentTwoAttachmentLinkPaperclipUnlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.75 11.5V3A2.5 2.5 0 0 0 8.25.5h-2.5A2.5 2.5 0 0 0 3.25 3v8.5a2 2 0 0 0 2 2h1a2 2 0 0 0 2-2V4a1 1 0 0 0-1-1h-.5a1 1 0 0 0-1 1v5.5"/>`),
		g.Group(children),
	)
}