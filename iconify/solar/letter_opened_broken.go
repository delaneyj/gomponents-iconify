package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterOpenedBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M22 14c0 3.771 0 5.657-1.172 6.828C19.657 22 17.771 22 14 22h-4c-3.771 0-5.657 0-6.828-1.172C2 19.657 2 17.771 2 14c0-3.771 0-5.657 1.172-6.828C3.825 6.518 4.7 6.229 6 6.102m12 0c1.3.127 2.175.416 2.828 1.07c.654.653.943 1.528 1.07 2.828M10 6h4m-3 3h2"/><path d="M14 2.002c1.707.014 2.647.11 3.268.73c.732.732.732 1.91.732 4.267v2.064c0 .46 0 .69-.094.892c-.095.202-.272.35-.626.644l-.72.6M10 2.002c-1.707.014-2.647.11-3.268.73C6 3.464 6 4.642 6 6.999v2.064c0 .46 0 .69.094.892c.095.202.272.35.626.644l1.439 1.2c1.837 1.53 2.755 2.295 3.841 2.295c.65 0 1.239-.273 2-.82"/></g>`),
		g.Group(children),
	)
}