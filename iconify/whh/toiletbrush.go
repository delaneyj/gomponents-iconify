package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toiletbrush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.36 960q-73 0-100.5-34t-27.5-94V160q6 8 19.5 29.5t19.5 28t19 18t29.5 14.5t40.5 3q70 0 128-35v709q-55 33-128 33zm-544 64q-96 0-177-17t-128-47t-47-64V224q55 42 148.5 67.5t203.5 25.5t203.5-25.5t148.5-67.5v672q0 34-47 64t-128 47t-177 17zm0-768q-96 0-177-17t-128-46.5t-47-64.5t47-64.5t128-46.5t177-17t177 17t128 46.5t47 64.5t-47 64.5t-128 46.5t-177 17zm0-192q-66 0-113 19t-47 45.5t47 45t113 18.5t113-18.5t47-45t-47-45.5t-113-19z"/>`),
		g.Group(children),
	)
}