package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AtFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 128v1.52c-.34 14.26-5.63 30.48-28 30.48c-23.14 0-28-17.4-28-32V88a8 8 0 0 0-8.53-8a8.17 8.17 0 0 0-7.47 8.27v4a48 48 0 1 0 6.73 64.05a40.19 40.19 0 0 0 3.38 5C175.48 168 185.71 176 204 176a54.81 54.81 0 0 0 9.22-.75a4 4 0 0 1 4.09 6a104.05 104.05 0 0 1-91.4 50.75c-54.78-1.1-99.71-45.14-101.83-99.89A104 104 0 1 1 232 128Zm-136 0a32 32 0 1 0 32-32a32 32 0 0 0-32 32Z"/>`),
		g.Group(children),
	)
}