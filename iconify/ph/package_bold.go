package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PackageBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m225.6 62.64l-88-48.17a19.91 19.91 0 0 0-19.2 0l-88 48.17A20 20 0 0 0 20 80.19v95.62a20 20 0 0 0 10.4 17.55l88 48.17a19.89 19.89 0 0 0 19.2 0l88-48.17a20 20 0 0 0 10.4-17.55V80.19a20 20 0 0 0-10.4-17.55ZM128 36.57L200 76l-21.43 11.73l-72-39.42Zm0 78.83L56 76l25.56-14l72 39.41ZM44 96.79l72 39.4v76.67l-72-39.42Zm96 116.07v-76.67l24-13.13V152a12 12 0 0 0 24 0v-42.08l24-13.13v76.65Z"/>`),
		g.Group(children),
	)
}