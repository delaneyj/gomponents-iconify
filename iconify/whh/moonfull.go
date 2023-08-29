package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moonfull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 0Q408 0 313 40.5t-163.5 109T40.5 313T0 512t40.5 199t109 163.5t163.5 109t199 40.5t199-40.5t163.5-109t109-163.5t40.5-199t-40.5-199t-109-163.5T711 40.5T512 0z"/>`),
		g.Group(children),
	)
}