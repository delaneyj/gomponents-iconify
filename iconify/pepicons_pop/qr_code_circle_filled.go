package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopQrCodeCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M9 9v2h2V9H9ZM8 7.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-.5-.5H8ZM15 9v2h2V9h-2Zm-1-1.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-.5-.5h-4ZM9 15v2h2v-2H9Zm-1-1.5a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5H8Z" clip-rule="evenodd"/><path d="M13.5 14a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5v-4Z"/><path fill-rule="evenodd" d="M5 6a1 1 0 0 1 1-1h3.5a1 1 0 0 1 0 2H6a1 1 0 0 1-1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6 5a1 1 0 0 1 1 1v3.5a1 1 0 0 1-2 0V6a1 1 0 0 1 1-1Zm0 16a1 1 0 0 1-1-1v-3.5a1 1 0 1 1 2 0V20a1 1 0 0 1-1 1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5 20a1 1 0 0 1 1-1h3.5a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1Zm16 0a1 1 0 0 1-1 1h-3.5a1 1 0 1 1 0-2H20a1 1 0 0 1 1 1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M20 21a1 1 0 0 1-1-1v-3.5a1 1 0 1 1 2 0V20a1 1 0 0 1-1 1Zm0-16a1 1 0 0 1 1 1v3.5a1 1 0 1 1-2 0V6a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M21 6a1 1 0 0 1-1 1h-3.5a1 1 0 1 1 0-2H20a1 1 0 0 1 1 1Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopQrCodeCircleFilled0)"/></g>`),
		g.Group(children),
	)
}