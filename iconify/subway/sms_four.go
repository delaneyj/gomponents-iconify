package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmsFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 0H64C28.6 0 0 28.6 0 64v256c0 35.4 28.6 64 64 64h128l-42.7 128l192-128H448c35.4 0 64-28.6 64-64V64c0-35.4-28.6-64-64-64zm-74.7 256L320 309.3l-64-64l-64 64l-53.3-53.3l64-64l-64-64L192 74.7l64 64l64-64l53.3 53.3l-64 64l64 64z"/>`),
		g.Group(children),
	)
}