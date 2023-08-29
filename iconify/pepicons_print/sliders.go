package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sliders(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path fill-rule="evenodd" d="M3 6.25a1 1 0 0 1 1-1h6.5a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1Zm11.375 0a1 1 0 0 1 1-1H17a1 1 0 1 1 0 2h-1.625a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path d="M16 6.5a2.75 2.75 0 1 1-5.5 0a2.75 2.75 0 0 1 5.5 0Z"/><path fill-rule="evenodd" d="M13.25 7.25a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm0 2a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5ZM3 16.25a1 1 0 0 1 1-1h6.5a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1Zm11.375 0a1 1 0 0 1 1-1H17a1 1 0 1 1 0 2h-1.625a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path d="M16 16.5a2.75 2.75 0 1 1-5.5 0a2.75 2.75 0 0 1 5.5 0Z"/><path fill-rule="evenodd" d="M13.25 17.25a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm0 2a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5ZM3 11.25a1 1 0 0 1 1-1h1.625a1 1 0 1 1 0 2H4a1 1 0 0 1-1-1Zm6.5 0a1 1 0 0 1 1-1H17a1 1 0 1 1 0 2h-6.5a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path d="M11 11.5a2.75 2.75 0 1 1-5.5 0a2.75 2.75 0 0 1 5.5 0Z"/><path fill-rule="evenodd" d="M8.25 12.25a.75.75 0 1 0 0-1.5a.75.75 0 0 0 0 1.5Zm0 2a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M3 4.75a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm11.38 0a.5.5 0 0 1 .5-.5h1.62a.5.5 0 0 1 0 1h-1.62a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.75 6.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM3 14.75a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm11.38 0a.5.5 0 0 1 .5-.5h1.62a.5.5 0 0 1 0 1h-1.62a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.75 16.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM3 9.75a.5.5 0 0 1 .5-.5h2.13a.5.5 0 0 1 0 1H3.5a.5.5 0 0 1-.5-.5Zm6.5 0a.5.5 0 0 1 .5-.5h6.5a.5.5 0 0 1 0 1H10a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M7.75 11.5a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3Zm0 1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}