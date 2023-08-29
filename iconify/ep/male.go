package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Male(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M399.5 849.5a225 225 0 1 0 0-450a225 225 0 0 0 0 450zm0 56.25a281.25 281.25 0 1 1 0-562.5a281.25 281.25 0 0 1 0 562.5zm253.125-787.5h225q28.125 0 28.125 28.125T877.625 174.5h-225q-28.125 0-28.125-28.125t28.125-28.125z"/><path fill="currentColor" d="M877.625 118.25q28.125 0 28.125 28.125v225q0 28.125-28.125 28.125T849.5 371.375v-225q0-28.125 28.125-28.125z"/><path fill="currentColor" d="M604.813 458.9L565.1 419.131l292.613-292.668l39.825 39.824z"/>`),
		g.Group(children),
	)
}