package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandMongodb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v19m6-10.773c0 3.273-1.812 4.77-6 9.273c-4.188-4.503-6-6-6-9.273C6 6.773 9.071 4.3 12 2c2.929 2.3 6 4.773 6 9.227z"/>`),
		g.Group(children),
	)
}