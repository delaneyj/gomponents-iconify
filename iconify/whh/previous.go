package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Previous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m962.428 1013l-563-463q-15-15-15-37.5t15-38.5l563-463q25-27 62 13v976q-37 40-62 13zm-834 11q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5t90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5z"/>`),
		g.Group(children),
	)
}