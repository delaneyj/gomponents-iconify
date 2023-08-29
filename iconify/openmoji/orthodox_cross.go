package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OrthodoxCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#b399c8" d="M12 12h48v47.83H12z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="3" d="M36 53V19m11 9.5H25m15.54-5.76h-9.08m9.28 26.89l-9.48-3.59"/><path stroke-width="2" d="M12 12h48v48H12z"/></g>`),
		g.Group(children),
	)
}