package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ballerina(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M162.007 250.666v69.585L220.499 512h61.832V296.68Zm120.324-61.977l-82.363 31.5l82.363 31.49zm0-44.99V0H162.007v189.723ZM120.314 250.665v69.585L61.81 512H0V296.68ZM0 188.689l82.342 31.5L0 251.68Zm0-44.99V0h120.314v189.723z"/>`),
		g.Group(children),
	)
}