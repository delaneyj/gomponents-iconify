package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesFourThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M80 44a36 36 0 1 0 36 36a36 36 0 0 0-36-36Zm0 64a28 28 0 1 1 28-28a28 28 0 0 1-28 28Zm96 8a36 36 0 1 0-36-36a36 36 0 0 0 36 36Zm0-64a28 28 0 1 1-28 28a28 28 0 0 1 28-28Zm-96 88a36 36 0 1 0 36 36a36 36 0 0 0-36-36Zm0 64a28 28 0 1 1 28-28a28 28 0 0 1-28 28Zm96-64a36 36 0 1 0 36 36a36 36 0 0 0-36-36Zm0 64a28 28 0 1 1 28-28a28 28 0 0 1-28 28Z"/>`),
		g.Group(children),
	)
}