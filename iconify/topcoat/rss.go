package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M40.5 41.5c0-22.092-17.908-40-40-40v8c17.673 0 32 14.326 32 32h8zm-16 0c0-13.256-10.745-24-24-24v7.998c8.836 0 16.001 7.166 16.001 16.002H24.5zm-16 0a8 8 0 0 0-8-8v8h8z"/>`),
		g.Group(children),
	)
}