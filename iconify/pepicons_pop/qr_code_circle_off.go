package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M9 9v2h2V9H9ZM8 7.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-.5-.5H8ZM15 9v2h2V9h-2Zm-1-1.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-.5-.5h-4ZM9 15v2h2v-2H9Zm-1-1.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5H8Z" clip-rule="evenodd"/><path d="M13.5 14a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5v-4Z"/><path fill-rule="evenodd" d="M5 6a1 1 0 0 1 1-1h3.5a1 1 0 0 1 0 2H6a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6 5a1 1 0 0 1 1 1v3.5a1 1 0 0 1-2 0V6a1 1 0 0 1 1-1Zm0 16a1 1 0 0 1-1-1v-3.5a1 1 0 1 1 2 0V20a1 1 0 0 1-1 1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5 20a1 1 0 0 1 1-1h3.5a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1Zm16 0a1 1 0 0 1-1 1h-3.5a1 1 0 1 1 0-2H20a1 1 0 0 1 1 1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M20 21a1 1 0 0 1-1-1v-3.5a1 1 0 1 1 2 0V20a1 1 0 0 1-1 1Zm0-16a1 1 0 0 1 1 1v3.5a1 1 0 1 1-2 0V6a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M21 6a1 1 0 0 1-1 1h-3.5a1 1 0 1 1 0-2H20a1 1 0 0 1 1 1Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}