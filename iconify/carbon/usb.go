package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 15V6a2 2 0 0 0-2-2H10a2 2 0 0 0-2 2v9a2 2 0 0 0-2 2v11h2V17h16v11h2V17a2 2 0 0 0-2-2ZM10 6h12v9H10Z"/><path fill="currentColor" d="M12 10h3v2h-3zm5 0h3v2h-3z"/>`),
		g.Group(children),
	)
}