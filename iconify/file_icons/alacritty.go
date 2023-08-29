package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alacritty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M175.09 330.675L224.098 512l49.007-181.325l-49.007-118.74l-49.008 118.74zM182.816 0h82.563l182.816 453.723H371.53l-147.432-346.54l-147.433 346.54H0L182.816 0z"/>`),
		g.Group(children),
	)
}