package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spinner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M304 48a48 48 0 1 0-96 0a48 48 0 1 0 96 0zm0 416a48 48 0 1 0-96 0a48 48 0 1 0 96 0zM48 304a48 48 0 1 0 0-96a48 48 0 1 0 0 96zm464-48a48 48 0 1 0-96 0a48 48 0 1 0 96 0zM142.9 437A48 48 0 1 0 75 369.1a48 48 0 1 0 67.9 67.9zm0-294.2A48 48 0 1 0 75 75a48 48 0 1 0 67.9 67.9zM369.1 437a48 48 0 1 0 67.9-67.9a48 48 0 1 0-67.9 67.9z"/>`),
		g.Group(children),
	)
}