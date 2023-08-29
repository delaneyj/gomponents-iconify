package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="m4.929 13.142l5.657 5.657a2 2 0 0 0 2.828 0l2.693-2.692L5.546 5.546a5 5 0 0 0-.617 7.596Z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.546 5.546a5 5 0 0 0-.617 7.596l5.657 5.657a2 2 0 0 0 2.828 0l2.693-2.692M10.89 5.232c.398.22.772.5 1.11.838a5 5 0 0 1 7.071 7.071l-.136.136"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m4 4l16 16"/></g>`),
		g.Group(children),
	)
}