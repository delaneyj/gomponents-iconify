package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShirtFoldedThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 52h-22.34l-14.83-14.83a4 4 0 0 0-2.44-1.17a2.13 2.13 0 0 0-.39 0H96a4 4 0 0 0-2.83 1.15L78.34 52H56a12 12 0 0 0-12 12v152a12 12 0 0 0 12 12h144a12 12 0 0 0 12-12V64a12 12 0 0 0-12-12Zm-28 5.66V112a4 4 0 0 1-2.3 3.63a3.93 3.93 0 0 1-4.21-.51l-32.2-27.82l27.33-41Zm-44 23.13L103.47 44h49.06ZM95.38 46.28l27.33 41l-32.17 27.8A4 4 0 0 1 84 112V57.66ZM52 216V64a4 4 0 0 1 4-4h20v52a11.89 11.89 0 0 0 6.91 10.87A12.08 12.08 0 0 0 88 124a11.88 11.88 0 0 0 7.65-2.8h.06L124 96.74V220H56a4 4 0 0 1-4-4Zm152 0a4 4 0 0 1-4 4h-68V96.74l28.32 24.46A11.9 11.9 0 0 0 168 124a12.08 12.08 0 0 0 5.12-1.15A11.89 11.89 0 0 0 180 112V60h20a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}