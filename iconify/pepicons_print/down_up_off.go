package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownUpOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g fill-rule="evenodd" clip-rule="evenodd" opacity=".2"><path d="M11.44 9.56a1.5 1.5 0 0 1 0-2.12l3-3a1.5 1.5 0 0 1 2.12 2.12l-3 3a1.5 1.5 0 0 1-2.12 0Z"/><path d="M19.56 9.56a1.5 1.5 0 0 1-2.12 0l-3-3a1.5 1.5 0 0 1 2.12-2.12l3 3a1.5 1.5 0 0 1 0 2.12Z"/><path d="M15.5 6A1.5 1.5 0 0 1 17 7.5v8a1.5 1.5 0 0 1-3 0v-8A1.5 1.5 0 0 1 15.5 6Zm-3.94 7.44a1.5 1.5 0 0 1 0 2.12l-3 3a1.5 1.5 0 0 1-2.12-2.12l3-3a1.5 1.5 0 0 1 2.12 0Z"/><path d="M3.44 13.44a1.5 1.5 0 0 1 2.12 0l3 3a1.5 1.5 0 0 1-2.12 2.12l-3-3a1.5 1.5 0 0 1 0-2.12Z"/><path d="M7.5 17A1.5 1.5 0 0 1 6 15.5v-8a1.5 1.5 0 1 1 3 0v8A1.5 1.5 0 0 1 7.5 17Z"/></g><path fill-rule="evenodd" d="M10.646 7.354a.5.5 0 0 1 0-.708l3-3a.5.5 0 0 1 .708.708l-3 3a.5.5 0 0 1-.708 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17.354 7.354a.5.5 0 0 1-.708 0l-3-3a.5.5 0 0 1 .708-.708l3 3a.5.5 0 0 1 0 .708Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M14 4a.5.5 0 0 1 .5.5V14a.5.5 0 0 1-1 0V4.5A.5.5 0 0 1 14 4Zm-4.646 8.646a.5.5 0 0 1 0 .708l-3 3a.5.5 0 0 1-.708-.708l3-3a.5.5 0 0 1 .708 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M2.646 12.646a.5.5 0 0 1 .708 0l3 3a.5.5 0 0 1-.708.708l-3-3a.5.5 0 0 1 0-.708Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6 16a.5.5 0 0 1-.5-.5V6a.5.5 0 0 1 1 0v9.5a.5.5 0 0 1-.5.5Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}