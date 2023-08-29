package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="flatColorIconsUpRight0" fill="#3F51B5" d="M44 19L30 30.7V7.3z"/><path id="flatColorIconsUpRight1" fill="#3F51B5" d="M6 27v13h8V27c0-2.2 1.8-4 4-4h17v-8H18c-6.6 0-12 5.4-12 12z"/></defs><use href="#flatColorIconsUpRight0"/><use href="#flatColorIconsUpRight1"/><use href="#flatColorIconsUpRight0"/><use href="#flatColorIconsUpRight1"/>`),
		g.Group(children),
	)
}