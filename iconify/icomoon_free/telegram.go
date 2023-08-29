package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Telegram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 0C3.581 0 0 3.581 0 8s3.581 8 8 8s8-3.581 8-8s-3.581-8-8-8zm3.931 5.484l-1.313 6.184c-.091.441-.356.544-.725.341l-2-1.478l-.959.934c-.112.109-.2.2-.4.2c-.259 0-.216-.097-.303-.344L5.55 9.084l-1.978-.616c-.428-.131-.431-.425.097-.634l7.706-2.975c.35-.159.691.084.556.625z"/>`),
		g.Group(children),
	)
}