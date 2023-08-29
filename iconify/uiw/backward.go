package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Backward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m10.903 2.931l-.001 5.108l6.615-5.593c.62-.526 1.58-.323 1.58.485v14.14c0 .805-.96 1.009-1.58.483l-6.615-5.593v5.11c0 .805-.96 1.009-1.58.483l-8.085-6.836a.936.936 0 0 1 0-1.434l8.086-6.838c.62-.526 1.58-.323 1.58.485Z"/>`),
		g.Group(children),
	)
}