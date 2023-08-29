package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m14.232 13.398l5.571-4.25a.5.5 0 0 0 0-.796l-5.571-4.25a.5.5 0 0 0-.803.398V13a.5.5 0 0 0 .803.398Zm.197-1.408V5.51l4.247 3.24l-4.247 3.24Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="m14.232 21.898l5.571-4.25a.5.5 0 0 0 0-.796l-5.571-4.25a.5.5 0 0 0-.803.398v8.5a.5.5 0 0 0 .803.398Zm.197-7.888l4.247 3.24l-4.247 3.24v-6.48Z" clip-rule="evenodd"/><path d="M14.227 12.6a.5.5 0 1 1-.597.8L6.665 8.207a.5.5 0 1 1 .598-.801l6.964 5.194Z"/><path d="M14.227 13.4a.5.5 0 1 0-.597-.8l-6.965 5.194a.5.5 0 0 0 .598.801l6.964-5.194Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}