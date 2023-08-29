package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="m15 14.797l-6.556 4.808a.75.75 0 0 1-.888-1.21l6.676-4.895l-6.676-4.895a.75.75 0 0 1 .888-1.21L15 12.203V4.5a.5.5 0 0 1 .8-.4l6 4.5a.5.5 0 0 1 0 .8l-5.467 4.1l5.467 4.1a.5.5 0 0 1 0 .8l-6 4.5a.5.5 0 0 1-.8-.4v-7.703Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="m14.232 13.398l5.571-4.25a.5.5 0 0 0 0-.796l-5.571-4.25a.5.5 0 0 0-.803.398V13a.5.5 0 0 0 .803.398Zm.197-1.408V5.51l4.247 3.24l-4.247 3.24Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="m14.232 21.898l5.571-4.25a.5.5 0 0 0 0-.796l-5.571-4.25a.5.5 0 0 0-.803.398v8.5a.5.5 0 0 0 .803.398Zm.197-7.888l4.247 3.24l-4.247 3.24v-6.48Z" clip-rule="evenodd"/><path d="M14.227 12.6a.5.5 0 1 1-.597.8L6.665 8.207a.5.5 0 1 1 .598-.801l6.964 5.194Z"/><path d="M14.227 13.4a.5.5 0 1 0-.597-.8l-6.965 5.194a.5.5 0 0 0 .598.801l6.964-5.194Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}