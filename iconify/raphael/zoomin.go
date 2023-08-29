package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zoomin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22.646 19.307a10.43 10.43 0 0 0 1.524-5.42c0-5.794-4.692-10.486-10.482-10.488c-5.79 0-10.484 4.693-10.484 10.485c0 5.79 4.693 10.48 10.484 10.48c1.987 0 3.84-.562 5.422-1.522l7.128 7.127l3.535-3.537l-7.127-7.126zm-8.958 1.062a6.495 6.495 0 0 1-6.484-6.485a6.494 6.494 0 0 1 6.484-6.486a6.493 6.493 0 0 1 6.484 6.485a6.496 6.496 0 0 1-6.484 6.484zm2-11.32h-4v2.834H8.853v4h2.833v2.834h4v-2.834h2.832v-4h-2.834V9.052z"/>`),
		g.Group(children),
	)
}