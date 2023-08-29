package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.5 5.5v3h3v-3h-3Zm-.5-1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V5a.5.5 0 0 0-.5-.5H5Zm6.5 1v3h3v-3h-3Zm-.5-1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V5a.5.5 0 0 0-.5-.5h-4Zm-5.5 7v3h3v-3h-3Zm-.5-1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5H5Z" clip-rule="evenodd"/><path d="M10.5 11a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5v-4Z"/><path fill-rule="evenodd" d="M11.5 11.5v3h3v-3h-3Zm-.5-1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-4ZM2.5 3a.5.5 0 0 1 .5-.5h3.5a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M3 2.5a.5.5 0 0 1 .5.5v3.5a.5.5 0 0 1-1 0V3a.5.5 0 0 1 .5-.5Zm0 15a.5.5 0 0 1-.5-.5v-3.5a.5.5 0 0 1 1 0V17a.5.5 0 0 1-.5.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M2.5 17a.5.5 0 0 1 .5-.5h3.5a.5.5 0 0 1 0 1H3a.5.5 0 0 1-.5-.5Zm15 0a.5.5 0 0 1-.5.5h-3.5a.5.5 0 0 1 0-1H17a.5.5 0 0 1 .5.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17 17.5a.5.5 0 0 1-.5-.5v-3.5a.5.5 0 0 1 1 0V17a.5.5 0 0 1-.5.5Zm0-15a.5.5 0 0 1 .5.5v3.5a.5.5 0 0 1-1 0V3a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17.5 3a.5.5 0 0 1-.5.5h-3.5a.5.5 0 0 1 0-1H17a.5.5 0 0 1 .5.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}