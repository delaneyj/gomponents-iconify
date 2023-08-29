package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0zm0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469zm0-405q-80 0-136 56T64 256t56 136t136 56t136-56t56-136t-56-136t-136-56zm0 341q-62 0-105.5-43.5T107 256t43.5-105.5T256 107t105.5 43.5T405 256t-43.5 105.5T256 405zm43-106h-22V171h-42l-30 23q-18 12-11 28q3 8 11.5 10.5T222 230l13-6v75h-22q-21 0-21 21t21 21h86q21 0 21-21t-21-21z"/>`),
		g.Group(children),
	)
}