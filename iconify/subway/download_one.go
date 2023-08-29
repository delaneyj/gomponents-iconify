package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M315.1 4.1v229.2h98.5L256 410.5L98.5 233.3H197V4.1C84.1 30.8 0 132 0 253c0 141.4 114.6 256 256 256s256-114.6 256-256c0-121-84.1-222.2-196.9-248.9z"/>`),
		g.Group(children),
	)
}