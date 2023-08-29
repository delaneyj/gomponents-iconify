package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Normalise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m256 0l256 256l-256 256L0 256L256 0zm-76.215 135.702v108.032L64.98 258.764l150.58 19.713V237.64l116.662 138.65V268.259l114.813-15.024l-150.58-19.712v40.814l-116.67-138.635z"/>`),
		g.Group(children),
	)
}