package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 206.8H315.1L256 9.8l-59.1 196.9H0L157.5 315l-59 187.2L256 403.7l157.5 98.5l-59.1-187.1L512 206.8zM256 354.5l-88.6 59.1l39.4-108.3l-78.8-59.1h98.5l29.5-98.5l29.5 98.5H384l-78.8 59.1l39.4 108.3l-88.6-59.1z"/>`),
		g.Group(children),
	)
}