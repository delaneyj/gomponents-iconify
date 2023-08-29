package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Faust(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M156.3 0h129.14l113.994 315.825H283.326l-66.602-202.596l-67.277 202.596h-109L156.302 0zm-31.53 449.615c0-47.846-52.151-77.913-93.648-53.99c-41.496 23.924-41.496 84.056 0 107.979c41.497 23.923 93.649-6.143 93.649-53.99zm317.214 0c0-47.846-52.152-77.913-93.649-53.99c-41.496 23.924-41.496 84.056 0 107.979c41.497 23.923 93.649-6.143 93.649-53.99z"/>`),
		g.Group(children),
	)
}