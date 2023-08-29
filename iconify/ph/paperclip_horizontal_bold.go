package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperclipHorizontalBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M252 128a60.07 60.07 0 0 1-60 60H44a40 40 0 0 1 0-80h140a12 12 0 0 1 0 24H44a16 16 0 0 0 0 32h148a36 36 0 0 0 0-72H80a12 12 0 0 1 0-24h112a60.07 60.07 0 0 1 60 60Z"/>`),
		g.Group(children),
	)
}