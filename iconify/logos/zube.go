package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#0D83DD" d="m0 63.228l119.833 63.23v50.744L0 113.972V63.227ZM119.834 0L8.994 58.483L57.08 83.855l62.754-33.111V0Zm16.333 0v50.744L207.914 88.6l-71.747 37.856V177.2L256 113.971V63.229L136.167 0Z"/>`),
		g.Group(children),
	)
}