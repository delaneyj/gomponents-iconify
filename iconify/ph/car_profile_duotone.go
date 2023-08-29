package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarProfileDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M88 176a24 24 0 1 1-24-24a24 24 0 0 1 24 24Zm104-24a24 24 0 1 0 24 24a24 24 0 0 0-24-24Zm-29.66-77.66a8 8 0 0 0-5.65-2.34H44.28a8 8 0 0 0-6.66 3.56L8 120h200Z" opacity=".2"/><path d="M240 112h-28.69L168 68.69A15.86 15.86 0 0 0 156.69 64H44.28A16 16 0 0 0 31 71.12L1.34 115.56A8.07 8.07 0 0 0 0 120v48a16 16 0 0 0 16 16h17a32 32 0 0 0 62 0h66a32 32 0 0 0 62 0h17a16 16 0 0 0 16-16v-40a16 16 0 0 0-16-16ZM44.28 80h112.41l32 32H23ZM64 192a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm128 0a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm48-24h-17a32 32 0 0 0-62 0H95a32 32 0 0 0-62 0H16v-40h224Z"/></g>`),
		g.Group(children),
	)
}