package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m11.232 10.398l5.571-4.25a.5.5 0 0 0 0-.796l-5.571-4.25a.5.5 0 0 0-.803.398V10a.5.5 0 0 0 .803.398Zm.197-1.408V2.51l4.247 3.24l-4.247 3.24Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="m11.232 18.898l5.571-4.25a.5.5 0 0 0 0-.796l-5.571-4.25a.5.5 0 0 0-.803.398v8.5a.5.5 0 0 0 .803.398Zm.197-7.888l4.247 3.24l-4.247 3.24v-6.48Z" clip-rule="evenodd"/><path d="M11.227 9.6a.5.5 0 1 1-.597.8L3.665 5.207a.5.5 0 1 1 .598-.801L11.227 9.6Z"/><path d="M11.227 10.4a.5.5 0 1 0-.597-.8l-6.965 5.194a.5.5 0 0 0 .598.801l6.964-5.194Z"/></g>`),
		g.Group(children),
	)
}