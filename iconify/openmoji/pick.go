package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#a57939" d="m37.5 11.5l-.1 31.1l1.2 3.6v15.7h-6V46.2l1.2-3.6V11.5h3.7z"/><path fill="#d0cfce" d="M58.6 23.3c-2-4.7-11.3-8.2-22.6-8.2c-11.4 0-20.8 3.6-22.7 8.3l8-1.5a80.71 80.71 0 0 1 29.4-.1Z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m37.5 20.4l-.1 21.9l1.2 3.7v15.7h-6V46l1.2-3.7V20.4h3.7zm0-8.7l-.1 3.2h-3.6v-3.2h3.7z"/><path d="M58.6 23.1c-2-4.7-11.3-8.2-22.6-8.2c-11.4 0-20.8 3.6-22.7 8.3l8-1.5a80.71 80.71 0 0 1 29.4-.1Z"/></g>`),
		g.Group(children),
	)
}