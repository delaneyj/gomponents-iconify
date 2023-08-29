package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilQrCodeCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M8.5 8.5v3h3v-3h-3Zm-.5-1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-.5-.5H8Zm6.5 1v3h3v-3h-3Zm-.5-1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-.5-.5h-4Zm-5.5 7v3h3v-3h-3Zm-.5-1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5H8Z" clip-rule="evenodd"/><path d="M13.5 14a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5v-4Z"/><path fill-rule="evenodd" d="M14.5 14.5v3h3v-3h-3Zm-.5-1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-4ZM5.5 6a.5.5 0 0 1 .5-.5h3.5a.5.5 0 0 1 0 1H6a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6 5.5a.5.5 0 0 1 .5.5v3.5a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5Zm0 15a.5.5 0 0 1-.5-.5v-3.5a.5.5 0 0 1 1 0V20a.5.5 0 0 1-.5.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5.5 20a.5.5 0 0 1 .5-.5h3.5a.5.5 0 0 1 0 1H6a.5.5 0 0 1-.5-.5Zm15 0a.5.5 0 0 1-.5.5h-3.5a.5.5 0 0 1 0-1H20a.5.5 0 0 1 .5.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M20 20.5a.5.5 0 0 1-.5-.5v-3.5a.5.5 0 0 1 1 0V20a.5.5 0 0 1-.5.5Zm0-15a.5.5 0 0 1 .5.5v3.5a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M20.5 6a.5.5 0 0 1-.5.5h-3.5a.5.5 0 0 1 0-1H20a.5.5 0 0 1 .5.5Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilQrCodeCircleFilled0)"/></g>`),
		g.Group(children),
	)
}