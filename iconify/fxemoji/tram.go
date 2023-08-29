package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#132028" d="m312 27.429l-11-11V16h-90v.429l-11 11L244.571 72L232 84.571L243.429 96L256 83.429L268.571 96L280 84.571L267.429 72z"/><path fill="#FF473E" d="M128 64h256v368H128z"/><path fill="#F9F9F7" d="M136 72h240v224H136z"/><path fill="#FFD469" d="M276 340c0 11.046-8.954 20-20 20s-20-8.954-20-20s8.954-20 20-20s20 8.954 20 20zm-97-4c-6.627 0-12 5.373-12 12s5.373 12 12 12s12-5.373 12-12s-5.373-12-12-12zm152 0c-6.627 0-12 5.373-12 12s5.373 12 12 12s12-5.373 12-12s-5.373-12-12-12z"/><path fill="#132028" d="M353 118H161V86h192v32zm-101 16H144v144h108V134zm8 0v144h108V134H260zm-51 362l16-64h-32l-64 64h80zm175 0l-64-64h-32l16 64h80z"/><path fill="#FFF" d="M227.429 32h57.142L256 60.571z"/>`),
		g.Group(children),
	)
}