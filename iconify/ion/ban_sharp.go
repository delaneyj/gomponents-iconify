package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BanSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M414.39 97.61A224 224 0 1 0 97.61 414.39A224 224 0 1 0 414.39 97.61ZM432 256a175.09 175.09 0 0 1-35.8 106.26L149.74 115.8A175.09 175.09 0 0 1 256 80c97.05 0 176 79 176 176Zm-352 0a175.09 175.09 0 0 1 35.8-106.26L362.26 396.2A175.09 175.09 0 0 1 256 432c-97 0-176-78.95-176-176Z"/>`),
		g.Group(children),
	)
}