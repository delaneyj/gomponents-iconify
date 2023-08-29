package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 207.9H315.1L256 11l-59.1 196.9H0l157.5 108.3l-59 187.1L256 404.8l157.5 98.5l-59-187.1z"/>`),
		g.Group(children),
	)
}