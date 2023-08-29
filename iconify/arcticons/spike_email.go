package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpikeEmail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.819 24a11.807 11.807 0 1 0-.672 3.937A11.819 11.819 0 0 0 35.82 24Zm1.903 16.552q.802-.665 1.537-1.406"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 24A21.5 21.5 0 1 0 5.802 35.412l-1.334 8.12l8.12-1.334a21.453 21.453 0 0 0 19.78 1.612"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.145 27.936c-1.444 3.522.907 5.417 2.93 5.877c2.467.56 7.425-2.019 7.425-9.813"/>`),
		g.Group(children),
	)
}