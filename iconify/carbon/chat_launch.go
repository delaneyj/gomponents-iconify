package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatLaunch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 4v2h4.586L20 12.586L21.414 14L28 7.414V12h2V4h-8zm6 12v4a1.996 1.996 0 0 1-2 2h-6l-4 7l1.736 1l3.429-6H26a4 4 0 0 0 4-4v-4zM4 20V8a1.996 1.996 0 0 1 2-2h12V4H6a3.999 3.999 0 0 0-4 4v12a4 4 0 0 0 4 4h9v-2H6a1.996 1.996 0 0 1-2-2z"/>`),
		g.Group(children),
	)
}