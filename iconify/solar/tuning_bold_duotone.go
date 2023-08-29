package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M13.752 14.5a3 3 0 1 1 6 0a3 3 0 0 1-6 0Zm-10-5a3 3 0 1 0 6 0a3 3 0 0 0-6 0Z"/><path d="M7.462 6.584a3.007 3.007 0 0 0-1.5.021V2a.75.75 0 0 1 1.5 0v4.584Zm-1.5 5.81a3.003 3.003 0 0 0 1.5.021V22a.75.75 0 0 1-1.5 0v-9.606Zm10 5V22a.75.75 0 0 0 1.5 0v-4.585a3.008 3.008 0 0 1-1.5-.021Zm1.5-5.81V2a.75.75 0 0 0-1.5 0v9.605a3.003 3.003 0 0 1 1.5-.021Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}