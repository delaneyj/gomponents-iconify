package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Digg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M328 26h204v983H0V312h328V26zm0 819V476H205v369h123zm286-533v697h205V312H614zm0-286v204h205V26H614zm287 286h533v942H901v-163h328v-82H901V312zm328 533V476h-123v369h123zm287-533h532v942h-532v-163h327v-82h-327V312zm327 533V476h-123v369h123z"/>`),
		g.Group(children),
	)
}