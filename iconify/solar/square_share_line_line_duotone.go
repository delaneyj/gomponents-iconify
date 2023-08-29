package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareShareLineLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M22 13.998c-.029 3.414-.218 5.296-1.46 6.537C19.076 22 16.718 22 12.003 22c-4.716 0-7.073 0-8.538-1.465C2 19.07 2 16.713 2 11.997C2 7.282 2 4.924 3.465 3.46C4.706 2.218 6.588 2.029 10.002 2" opacity=".5"/><path stroke-linejoin="round" d="M22 7h-8c-1.818 0-2.913.892-3.32 1.3c-.125.127-.188.19-.19.19c0 .002-.063.065-.19.19C9.892 9.087 9 10.182 9 12v3m13-8l-5-5m5 5l-5 5"/></g>`),
		g.Group(children),
	)
}