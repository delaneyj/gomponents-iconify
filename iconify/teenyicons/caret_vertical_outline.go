package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretVerticalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m7.5 13l-.332.374l.332.295l.332-.295L7.5 13Zm0-11l.34-.367l-.333-.308l-.34.301L7.5 2Zm.332 11.374l4.5-4l-.664-.748l-4.5 4l.664.748Zm0-.748l-4.5-4l-.664.748l4.5 4l.664-.748Zm-.664-11l-4.5 4l.664.748l4.5-4l-.664-.748Zm-.008.74l4.313 4l.68-.733l-4.313-4l-.68.734Z"/>`),
		g.Group(children),
	)
}