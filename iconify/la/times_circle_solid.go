package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TimesCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 3C8.832 3 3 8.832 3 16s5.832 13 13 13s13-5.832 13-13S23.168 3 16 3zm0 2c6.086 0 11 4.914 11 11s-4.914 11-11 11S5 22.086 5 16S9.914 5 16 5zm-3.781 5.781L10.78 12.22L14.562 16l-3.78 3.781l1.437 1.438L16 17.437l3.781 3.782l1.438-1.438L17.437 16l3.782-3.781l-1.438-1.438L16 14.562z"/>`),
		g.Group(children),
	)
}