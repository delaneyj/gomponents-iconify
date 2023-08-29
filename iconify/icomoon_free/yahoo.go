package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yahoo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8.878 9.203C10.637 6.115 13.544 1.078 14.341 0c-.35.234-.887.353-1.381.466L12.213 0c-.6 1.119-2.813 4.734-4.222 7.05C6.563 4.684 4.872 1.953 3.769 0C2.894.188 2.532.197 1.66 0c1.731 2.606 4.503 7.572 5.447 9.203L6.979 16l1.013-.466v-.012L9.004 16l-.125-6.797z"/>`),
		g.Group(children),
	)
}