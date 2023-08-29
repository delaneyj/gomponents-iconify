package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 4.06v-.5a.509.509 0 0 0-.15-.35a.483.483 0 0 0-.7 0a.491.491 0 0 0-.15.35v.5a8.41 8.41 0 0 0-7.88 7.82a.976.976 0 0 0 .27.74a1.029 1.029 0 0 0 .74.32h6.87v5.22a1.653 1.653 0 0 1-.62 1.54a1.528 1.528 0 0 1-2.38-1.16a.5.5 0 0 0-1 0a2.433 2.433 0 0 0 2.43 2.4a2.45 2.45 0 0 0 2.57-2.29c.08-1.39 0-2.81 0-4.2v-1.51h6.87a1.029 1.029 0 0 0 .74-.32a.976.976 0 0 0 .27-.74a8.41 8.41 0 0 0-7.88-7.82Zm6.87 7.88l-14.75.01a7.4 7.4 0 0 1 14.76-.02c0 .01 0 .01-.01.01Z"/>`),
		g.Group(children),
	)
}