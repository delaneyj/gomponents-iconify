package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardTransferLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M22 11c-.01-3.114-.108-4.765-1.172-5.828C19.657 4 17.771 4 14 4h-4C6.229 4 4.343 4 3.172 5.172C2 6.343 2 8.229 2 12c0 3.771 0 5.657 1.172 6.828C4.343 20 6.229 20 10 20h1.5"/><path stroke-linejoin="round" d="M20 20v-6m0 0l2 2m-2-2l-2 2" opacity=".5"/><path stroke-linejoin="round" d="M15.5 14v6m0 0l2-2m-2 2l-2-2"/><path d="M10 16H6" opacity=".5"/><path d="M2 10h20" opacity=".4"/></g>`),
		g.Group(children),
	)
}