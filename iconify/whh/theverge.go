package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Theverge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M576 896H448L0 128L64 0h896l64 128zM320 192l192 320l192-320H320z"/>`),
		g.Group(children),
	)
}