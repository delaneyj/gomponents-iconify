package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdhesivePlasterTwoLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m20.416 12.765l-9.181-9.18a5.41 5.41 0 0 0-7.65 7.65l9.18 9.18a5.41 5.41 0 1 0 7.65-7.65Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m12.765 20.416l7.65-7.65m-9.18-9.182l-7.65 7.651" opacity=".5"/><circle cx="9.172" cy="12" r="1" fill="currentColor" transform="rotate(-45 9.172 12)"/><circle cx="12" cy="14.829" r="1" fill="currentColor" transform="rotate(-45 12 14.829)"/><circle cx="12" cy="9.171" r="1" fill="currentColor" transform="rotate(-45 12 9.171)"/><circle cx="14.829" cy="12" r="1" fill="currentColor" transform="rotate(-45 14.829 12)"/></g>`),
		g.Group(children),
	)
}