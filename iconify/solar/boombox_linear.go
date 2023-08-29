package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoomboxLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M6 9.5h12"/><path d="M2 14c0-3.771 0-5.657 1.172-6.828C4.343 6 6.229 6 10 6h4c3.771 0 5.657 0 6.828 1.172C22 8.343 22 10.229 22 14c0 3.771 0 5.657-1.172 6.828C19.657 22 17.771 22 14 22h-4c-3.771 0-5.657 0-6.828-1.172C2 19.657 2 17.771 2 14Zm14-8c0-1.886 0-2.828-.586-3.414C14.828 2 13.886 2 12 2c-1.886 0-2.828 0-3.414.586C8 3.172 8 4.114 8 6"/><path d="M11 16a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm7.121 2.121a3 3 0 1 1-4.242-4.242a3 3 0 0 1 4.242 4.242Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M8 13v2m10.121-1.121l-1.414 1.414"/></g>`),
		g.Group(children),
	)
}