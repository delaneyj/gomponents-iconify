package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineEndArrowRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.525 18.025q-.5.325-1.012.038T11 17.175V13H3q-.425 0-.713-.288T2 12q0-.425.288-.713T3 11h8V6.825q0-.6.513-.888t1.012.038l8.15 5.175q.475.3.475.85t-.475.85l-8.15 5.175Z"/>`),
		g.Group(children),
	)
}