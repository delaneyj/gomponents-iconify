package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6 6v2h2V6H6ZM5 4.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V5a.5.5 0 0 0-.5-.5H5ZM12 6v2h2V6h-2Zm-1-1.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V5a.5.5 0 0 0-.5-.5h-4ZM6 12v2h2v-2H6Zm-1-1.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5H5Z" clip-rule="evenodd"/><path d="M10.5 11a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5v-4Z"/><path fill-rule="evenodd" d="M2 3a1 1 0 0 1 1-1h3.5a1 1 0 0 1 0 2H3a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M3 2a1 1 0 0 1 1 1v3.5a1 1 0 0 1-2 0V3a1 1 0 0 1 1-1Zm0 16a1 1 0 0 1-1-1v-3.5a1 1 0 1 1 2 0V17a1 1 0 0 1-1 1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M2 17a1 1 0 0 1 1-1h3.5a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1Zm16 0a1 1 0 0 1-1 1h-3.5a1 1 0 1 1 0-2H17a1 1 0 0 1 1 1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17 18a1 1 0 0 1-1-1v-3.5a1 1 0 1 1 2 0V17a1 1 0 0 1-1 1Zm0-16a1 1 0 0 1 1 1v3.5a1 1 0 1 1-2 0V3a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M18 3a1 1 0 0 1-1 1h-3.5a1 1 0 1 1 0-2H17a1 1 0 0 1 1 1Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}