package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folderalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.232 896h-896q-26 0-45-19t-19-45V256h1024v576q0 26-19 45t-45 19zm-960-704V64q0-27 18.5-45.5T64.232 0h384q26 0 45 18.5t19 45.5t18.5 45.5t45.5 18.5h384q26 0 45 19t19 45H.232z"/>`),
		g.Group(children),
	)
}