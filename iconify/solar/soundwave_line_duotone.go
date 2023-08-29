package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundwaveLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M12 4v16"/><path d="M16 7v10M8 7v10" opacity=".5"/><path d="M20 11v2M4 11v2"/></g>`),
		g.Group(children),
	)
}