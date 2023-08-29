package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pingpong(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M865 1020L602 713q-68 44-145.5 53.5t-157.5-20T152 650Q72 570 31.5 463t-29-211t78-171t171-78.5t211.5 29T649 152q68 67 97.5 147t20 157.5T714 602l308 262q10 10-23.5 55T922 999t-57 21zM605 190q-66-66-155-100T274 66t-142.5 65.5T66 274t24 176t100 155q64 64 142 87l360-360q-23-78-87-142z"/>`),
		g.Group(children),
	)
}