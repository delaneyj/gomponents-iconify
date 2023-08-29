package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lufthansa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.513 32.452C20.506 30.482 15.393 21.399 6.5 19.62"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.5 20.002a106.672 106.672 0 0 0-13.72-.382c-.094 1.904-4.87 5.44-8.025 7.103"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.573 21.937a91.145 91.145 0 0 0-12.882-.359m11.506 2.212a106.39 106.39 0 0 0-13.424-.412m11.052 2.204c-8.067-.699-13.206-.6-13.206-.6m-11.482-3.711l.635-1.047s4.605 1.842 6.575 2.731"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}