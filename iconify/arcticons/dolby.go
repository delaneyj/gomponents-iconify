package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dolby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.182 11.59H43.5v24.82h0h-5.318A12.41 12.41 0 0 1 25.772 24v0a12.41 12.41 0 0 1 12.41-12.41ZM9.819 36.41H4.5h0V11.59h5.318A12.41 12.41 0 0 1 22.228 24h0a12.41 12.41 0 0 1-12.41 12.41Z"/>`),
		g.Group(children),
	)
}