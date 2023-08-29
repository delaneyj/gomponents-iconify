package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StoolLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M198 64V40a14 14 0 0 0-14-14H72a14 14 0 0 0-14 14v24a14 14 0 0 0 14 14h9L58.07 223.06a6 6 0 0 0 5 6.87a6 6 0 0 0 6.86-5L78 174h100l8 50.93a6 6 0 1 0 11.86-1.87L175 78h9a14 14 0 0 0 14-14ZM70 64V40a2 2 0 0 1 2-2h112a2 2 0 0 1 2 2v24a2 2 0 0 1-2 2H72a2 2 0 0 1-2-2Zm106.14 98H79.86l13.27-84h69.74Z"/>`),
		g.Group(children),
	)
}