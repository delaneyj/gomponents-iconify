package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shiftcal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" d="M5.5 13.136h37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.36 16.518h5.16v5.16H9.36zm0 7.997h5.16v5.16H9.36zm0 7.997h5.16v5.16H9.36zm24.12-15.994h5.16v5.16h-5.16zm0 7.997h5.16v5.16h-5.16zm0 7.997h5.16v5.16h-5.16zm-8.04-15.994h5.16v5.16h-5.16zm0 7.997h5.16v5.16h-5.16zm0 7.997h5.16v5.16h-5.16zM17.4 16.518h5.16v5.16H17.4zm0 7.997h5.16v5.16H17.4zm0 7.997h5.16v5.16H17.4z"/>`),
		g.Group(children),
	)
}