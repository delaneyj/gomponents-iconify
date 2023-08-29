package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rootsbedrock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M.4 0L0 .4v5.2l.343.343l11.314-1.886L12 4.4V8l11.52-1.92l.48-.48V.4l-.4-.4zm.08 9.92L0 10.4v3.2l.343.343L12 12V8zM12 12v4l11.52-1.92l.48-.48v-3.2l-.343-.343zM.48 17.92L0 18.4v5.2l.4.4h23.2l.4-.4v-5.2l-.343-.343l-11.314 1.886L12 19.6V16L.48 17.92z"/>`),
		g.Group(children),
	)
}