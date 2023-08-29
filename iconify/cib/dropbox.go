package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 2.412L0 7.505l8 5.093l8-5.093zm16 0l-8 5.093l8 5.093l8-5.093zM0 17.697l8 5.1l8-5.1l-8-5.093zm24-5.093l-8 5.093l8 5.1l8-5.1zM8 24.495l8 5.093l8-5.093l-8-5.093z"/>`),
		g.Group(children),
	)
}