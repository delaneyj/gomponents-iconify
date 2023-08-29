package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PricetagAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 6a1 1 0 1 0 1 1a1 1 0 0 0-1-1Zm14.71 5.78l-9.48-9.46A1 1 0 0 0 11.5 2h-6a1 1 0 0 0-.71.29l-2.5 2.49a1 1 0 0 0-.29.71v6a1.05 1.05 0 0 0 .29.71l9.49 9.5a1.05 1.05 0 0 0 .71.29a1 1 0 0 0 .71-.29l8.51-8.51a1 1 0 0 0 .29-.71a1.05 1.05 0 0 0-.29-.7Zm-9.22 7.81L4 11.09V5.9L5.9 4h5.18l8.5 8.49Z"/>`),
		g.Group(children),
	)
}