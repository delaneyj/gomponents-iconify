package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Asterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m28.316 13.949l-.632-1.898L17 15.612V4h-2v11.612L4.316 12.051l-.632 1.898l10.684 3.561L7.2 27.066l1.6 1.201l7.2-9.6l7.2 9.6l1.6-1.201l-7.168-9.556l10.684-3.561z"/>`),
		g.Group(children),
	)
}