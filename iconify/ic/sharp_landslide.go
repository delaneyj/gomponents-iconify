package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpLandslide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.47 13.79l-2.58-1.03L6 15.05l-4-1.54v2.1l4 1.34zm-4.9-2.37L8 8H2v3.61l4 1.34zM6 19.05l-4-1.33V22h20l-4.97-6.62zM17 6V1l-5-1l-3 2v4l3 2zm1.5 1L16 9v3l2.5 2l4.5-2V8z"/>`),
		g.Group(children),
	)
}