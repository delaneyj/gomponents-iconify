package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCrossRoundedLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="6" r="4"/><path d="M14 20.834c-.634.108-1.305.166-2 .166c-3.866 0-7-1.79-7-4s3.134-4 7-4c1.713 0 3.283.352 4.5.936" opacity=".5"/><circle cx="17" cy="18" r="4"/><path stroke-linecap="round" stroke-linejoin="round" d="m15.666 16.667l2.667 2.666m0-2.666l-2.666 2.666"/></g>`),
		g.Group(children),
	)
}