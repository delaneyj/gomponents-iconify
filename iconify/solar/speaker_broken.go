package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M20 12v2c0 3.771 0 5.657-1.172 6.828C17.657 22 15.771 22 12 22c-3.771 0-5.657 0-6.828-1.172C4 19.657 4 17.771 4 14v-4c0-3.771 0-5.657 1.172-6.828C6.343 2 8.229 2 12 2c3.771 0 5.657 0 6.828 1.172c.944.943 1.127 2.348 1.163 4.828"/><path d="M14 7.5a2 2 0 1 1-4 0a2 2 0 0 1 4 0Zm1 8a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/></g>`),
		g.Group(children),
	)
}