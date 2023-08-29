package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.199 2.587a1.65 1.65 0 0 1 1.602 0l7.2 4c.524.291.849.843.849 1.443v7.94c0 .6-.325 1.152-.849 1.443l-7.2 4a1.65 1.65 0 0 1-1.602 0l-7.2-4a1.65 1.65 0 0 1-.849-1.443V8.03c0-.6.325-1.152.849-1.443l7.2-4Zm.874 1.312a.15.15 0 0 0-.146 0l-7.2 4a.15.15 0 0 0-.077.13v7.941c0 .055.03.105.077.132l7.2 4a.15.15 0 0 0 .146 0l7.2-4a.15.15 0 0 0 .077-.132V8.03a.15.15 0 0 0-.077-.131l-7.2-4Z"/><path d="M7.25 12a4.75 4.75 0 1 1 9.5 0a4.75 4.75 0 0 1-9.5 0ZM12 8.75a3.25 3.25 0 1 0 0 6.5a3.25 3.25 0 0 0 0-6.5Z"/></g>`),
		g.Group(children),
	)
}