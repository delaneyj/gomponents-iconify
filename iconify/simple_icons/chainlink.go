package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chainlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0L9.798 1.266l-6 3.468L1.596 6v12l2.202 1.266l6.055 3.468L12.055 24l2.202-1.266l5.945-3.468L22.404 18V6l-2.202-1.266l-6-3.468zM6 15.468V8.532l6-3.468l6 3.468v6.936l-6 3.468z"/>`),
		g.Group(children),
	)
}