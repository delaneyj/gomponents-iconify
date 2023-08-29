package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardSendBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M19 20v-6m0 0l2 2m-2-2l-2 2"/><path d="M22 12c0-3.771 0-5.657-1.172-6.828C19.657 4 17.771 4 14 4m0 16h-4c-3.771 0-5.657 0-6.828-1.172C2 17.657 2 15.771 2 12c0-3.771 0-5.657 1.172-6.828C4.343 4 6.229 4 10 4m0 12H6m7 0h-.5M2 10h5m15 0H11"/></g>`),
		g.Group(children),
	)
}