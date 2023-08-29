package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MouseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.141 2h1.718c2.014 0 3.094.278 4.071.801A5.452 5.452 0 0 1 19.2 5.07c.522.978.801 2.058.801 4.072v5.718c0 2.014-.279 3.094-.801 4.071a5.452 5.452 0 0 1-2.27 2.269c-.977.522-2.057.801-4.071.801H11.14c-2.014 0-3.094-.279-4.072-.801A5.452 5.452 0 0 1 4.8 18.931c-.522-.978-.8-2.058-.8-4.071V9.14c0-2.014.278-3.094.801-4.072A5.452 5.452 0 0 1 7.07 2.801C8.047 2.278 9.127 2 11.141 2ZM11 6v5h2V6h-2Z"/>`),
		g.Group(children),
	)
}