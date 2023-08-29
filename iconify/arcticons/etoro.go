package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Etoro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.459 36.447a7.6 7.6 0 0 1-6.606 3.836h0a7.602 7.602 0 0 1-7.602-7.602v-4.942a7.602 7.602 0 0 1 7.602-7.602h0a7.602 7.602 0 0 1 7.603 7.602v2.47H16.25m.294-10.072c-1.952 1.992-3.196 4.298-3.196 8.98c-3.56-.404-7.848-2.305-7.848-6.796S12.944 9.375 15.533 7.717C13.429 11.722 12.7 15.282 12.7 17.75s3.843 2.387 3.843 2.387Zm14.912 0c1.952 1.992 3.196 4.298 3.196 8.98c3.56-.404 7.848-2.305 7.848-6.796c0-4.49-7.444-12.946-10.033-14.604c2.104 4.005 2.832 7.565 2.832 10.033s-3.843 2.387-3.843 2.387Z"/>`),
		g.Group(children),
	)
}