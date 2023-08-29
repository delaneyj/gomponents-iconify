package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Iphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M76 0h265c42 0 76 34 76 76v616c0 42-34 76-76 76H76c-42 0-76-34-76-76V76C0 34 34 0 76 0zM49 617h319V144H49v473zm159 105c24 0 42-18 42-41c0-24-18-42-42-42c-23 0-41 18-41 42c0 23 18 41 41 41z"/>`),
		g.Group(children),
	)
}