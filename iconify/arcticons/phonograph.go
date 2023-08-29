package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phonograph(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.63 25L18.24 4.64a.48.48 0 0 0-.82.31l-.24 3.4a19 19 0 0 1-.95 4.65L9.77 31.92h0a8.56 8.56 0 1 0 14.3-2.78L30.26 27a19.09 19.09 0 0 1 4.67-.95l3.39-.24a.48.48 0 0 0 .31-.81ZM17.78 37.26a2.32 2.32 0 1 1 2.33-2.32a2.32 2.32 0 0 1-2.33 2.32Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.08 29.15a46.18 46.18 0 0 0-7.65 3.9"/>`),
		g.Group(children),
	)
}