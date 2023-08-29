package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandSnapseed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8.152 3.115a.46.46 0 0 0-.609 0C4.6 5.695 3.014 8.556 3 11.493c0 2.928 1.586 5.803 4.543 8.392a.46.46 0 0 0 .61 0c2.957-2.589 4.547-5.464 4.547-8.392c0-2.928-1.6-5.799-4.548-8.378z"/><path d="m8 20l12.09-.011c.503 0 .91-.434.91-.969v-6.063c0-.535-.407-.968-.91-.968h-7.382"/></g>`),
		g.Group(children),
	)
}