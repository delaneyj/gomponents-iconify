package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TuningThreeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M12 8.75a.75.75 0 0 1-.75-.75V5a.75.75 0 0 1 1.5 0v3a.75.75 0 0 1-.75.75Z"/><path fill-rule="evenodd" d="M4 7.25a2.75 2.75 0 1 1 0 5.5a2.75 2.75 0 0 1 0-5.5ZM5.25 10a1.25 1.25 0 1 0-2.5 0a1.25 1.25 0 0 0 2.5 0Zm9.5 2a2.75 2.75 0 1 1-5.5 0a2.75 2.75 0 0 1 5.5 0ZM12 13.25a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Zm10.75.75a2.75 2.75 0 1 1-5.5 0a2.75 2.75 0 0 1 5.5 0ZM20 15.25a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5Z" clip-rule="evenodd"/><path d="M19.25 10a.75.75 0 0 0 1.5 0V5a.75.75 0 0 0-1.5 0v5ZM4 13.25a.75.75 0 0 0-.75.75v5a.75.75 0 0 0 1.5 0v-5a.75.75 0 0 0-.75-.75ZM11.25 19a.75.75 0 0 0 1.5 0v-3a.75.75 0 0 0-1.5 0v3Zm8.75.75a.75.75 0 0 1-.75-.75v-1a.75.75 0 0 1 1.5 0v1a.75.75 0 0 1-.75.75ZM3.25 5a.75.75 0 0 1 1.5 0v1a.75.75 0 0 1-1.5 0V5Z"/></g>`),
		g.Group(children),
	)
}