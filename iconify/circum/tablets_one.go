package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletsOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.76 3.065a6.171 6.171 0 0 0-6.11 5.58a6.159 6.159 0 1 0 6.71 6.71a6.159 6.159 0 0 0-.6-12.29Zm-5.53 16.87A5.166 5.166 0 0 1 5.24 11.5l7.27 7.26a5.153 5.153 0 0 1-3.28 1.175Zm3.99-1.88l-7.27-7.27a5.165 5.165 0 0 1 7.27 7.27Zm2.15-3.71a6.12 6.12 0 0 0-5.72-5.71a5.157 5.157 0 1 1 5.72 5.71Z"/>`),
		g.Group(children),
	)
}