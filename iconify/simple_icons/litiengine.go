package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Litiengine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m0 12.018l2.09 2.088L11.987 24l2.146-2.146l-9.897-9.893l6.586-6.582l-2.09-2.089Zm13.211 6.624l2.08 2.078l5.425-5.422l-2.08-2.078zM9.85 2.151l6.606 6.602L9.9 15.306l2.134 2.133l6.555-6.553l3.258 3.257L24 11.993L12 0Zm-3.276 9.853l2.035 2.034l5.453-5.45l-2.035-2.035z"/>`),
		g.Group(children),
	)
}