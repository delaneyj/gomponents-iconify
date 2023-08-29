package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bomb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 1h1v1h-1V1zm0 4h1v1h-1V5zm2-2h1v1h-1V3zm-4 0h1v1h-1V3zm4.6-.9l.7-.7l-.7-.7l-1.4 1.4l.7.7zm-.7 2.1l-.7.7l1.4 1.4l.7-.7l-.7-.7zm-2.8-1.4l.7-.7L10.4.7l-.7.7l.7.7z"/><path fill="currentColor" d="m10.4 6.4l2-2l-.7-.7l-2 2L9 5l-.7.8C7.5 5.3 6.5 5 5.5 5C2.5 5 0 7.5 0 10.5S2.5 16 5.5 16s5.5-2.5 5.5-5.5c0-1-.3-1.9-.7-2.8L11 7l-.6-.6zM6 7.2C4 7.2 2.6 9 2.6 10h-1C1.6 8 4 6.2 6 6.2v1z"/>`),
		g.Group(children),
	)
}