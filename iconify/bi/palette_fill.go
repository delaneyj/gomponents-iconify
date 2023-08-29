package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaletteFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.433 10.07C14.133 10.585 16 11.15 16 8a8 8 0 1 0-8 8c1.996 0 1.826-1.504 1.649-3.08c-.124-1.101-.252-2.237.351-2.92c.465-.527 1.42-.237 2.433.07zM8 5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3zm4.5 3a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3zM5 6.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0zm.5 6.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3z"/>`),
		g.Group(children),
	)
}