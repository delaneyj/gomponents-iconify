package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pigpenj(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M928 1024H96q-40 0-68-28T0 928t28-68t68-28h736V96q0-40 28-68t68-28t68 28t28 68v832q0 40-28 68t-68 28zM512 640q-53 0-90.5-37.5T384 512t37.5-90.5T512 384t90.5 37.5T640 512t-37.5 90.5T512 640z"/>`),
		g.Group(children),
	)
}