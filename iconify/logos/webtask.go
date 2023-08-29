package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webtask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#5F3237" d="M256 256H0V0h256z"/><path fill="#E56E62" d="M256 256H51.2V51.2H256z"/><path fill="#ED955B" d="M204.8 204.8H51.2V51.2h153.6z"/><path fill="#F9E862" d="M204.8 204.8H102.4V102.4h102.4z"/><path fill="#FFF" d="M153.6 153.6h-51.2v-51.2h51.2z"/>`),
		g.Group(children),
	)
}