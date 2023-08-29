package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reverseimagesearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.075 17.927A16.765 16.765 0 1 1 26.608 38.46m-10.073-.44a16.765 16.765 0 0 1-10.46-11.415"/><circle cx="22.269" cy="22.382" r="1.734" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="29.785" cy="22.382" r="1.734" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="14.754" cy="22.382" r="1.734" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.972 34.273l7.953 8.227"/>`),
		g.Group(children),
	)
}