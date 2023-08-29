package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserBlockRoundedLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="6" r="4"/><circle cx="17" cy="18" r="4"/><path d="m20 15l-6 6m0-.166c-.634.108-1.305.166-2 .166c-3.866 0-7-1.79-7-4s3.134-4 7-4c1.713 0 3.283.352 4.5.936"/></g>`),
		g.Group(children),
	)
}