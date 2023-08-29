package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShuffleAngularDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M90.39 75.35L128 128l-37.61 52.65a8 8 0 0 1-6.51 3.35H32V72h51.88a8 8 0 0 1 6.51 3.35ZM172.12 72a8 8 0 0 0-6.51 3.35L128 128l37.61 52.65a8 8 0 0 0 6.51 3.35H232V72Z" opacity=".2"/><path d="M237.66 178.34a8 8 0 0 1 0 11.32l-24 24a8 8 0 0 1-11.32-11.32L212.69 192h-40.57a16 16 0 0 1-13-6.7L83.88 80H32a8 8 0 0 1 0-16h51.88a16 16 0 0 1 13 6.7L172.12 176h40.57l-10.35-10.34a8 8 0 0 1 11.32-11.32ZM143 107a8 8 0 0 0 11.16-1.86l18-25.12h40.57l-10.39 10.32a8 8 0 0 0 11.32 11.32l24-24a8 8 0 0 0 0-11.32l-24-24a8 8 0 0 0-11.32 11.32L212.69 64h-40.57a16 16 0 0 0-13 6.7l-17.97 25.12A8 8 0 0 0 143 107Zm-30 42a8 8 0 0 0-11.16 1.86L83.88 176H32a8 8 0 0 0 0 16h51.88a16 16 0 0 0 13-6.7l17.95-25.12A8 8 0 0 0 113 149Z"/></g>`),
		g.Group(children),
	)
}