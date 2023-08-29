package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BezierCurveThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M217.83 148.07A92.6 92.6 0 0 0 165.53 84H240a4 4 0 0 0 0-8h-84.29a28 28 0 0 0-55.42 0H16a4 4 0 0 0 0 8h74.47a92.6 92.6 0 0 0-52.3 64.07a28 28 0 1 0 8.07.64a84.51 84.51 0 0 1 55-60.36a28 28 0 0 0 53.46 0a84.53 84.53 0 0 1 55 60.36a28 28 0 1 0 8.07-.64ZM60 176a20 20 0 1 1-20-20a20 20 0 0 1 20 20Zm68-76a20 20 0 1 1 20-20a20 20 0 0 1-20 20Zm88 96a20 20 0 1 1 20-20a20 20 0 0 1-20 20Z"/>`),
		g.Group(children),
	)
}