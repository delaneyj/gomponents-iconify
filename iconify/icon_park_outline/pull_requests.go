package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PullRequests(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M37 44a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM11 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 32a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path stroke-linecap="round" d="M11 12v24m13-26h9a4 4 0 0 1 4 4v22"/><path stroke-linecap="round" d="m30 16l-6-6l6-6"/></g>`),
		g.Group(children),
	)
}