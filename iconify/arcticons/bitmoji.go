package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitmoji(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.59 5.5H12.41a4.57 4.57 0 0 0-4.56 4.57v21.25a4.58 4.58 0 0 0 4.56 4.57h10v6.61L29 35.89h6.56a4.58 4.58 0 0 0 4.56-4.57V10.07a4.57 4.57 0 0 0-4.53-4.57Z"/><rect width="4.98" height="6.33" x="26.35" y="12.59" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.49"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.16 13.26h0a2.51 2.51 0 0 0-2.49 2.5v.67c0 1.37 1.12-.24 2.49-.24h0c1.37 0 2.49 1.61 2.49.24v-.67a2.51 2.51 0 0 0-2.49-2.5Zm-3.75 8.95a8.59 8.59 0 1 0 17.18 0Z"/>`),
		g.Group(children),
	)
}