package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandCakephp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 11l8 2c1.361-.545 2-1.248 2-2V7.2C22 5.435 17.521 4 11.998 4C6.476 4 2 5.435 2 7.2V10c0 1.766 4.478 4 10 4v-3zm0 3v3l8 2c1.362-.547 2-1.246 2-2v-3c0 .754-.638 1.453-2 2l-8-2zM2 17c0 1.766 4.476 3 9.998 3L12 17c-5.522 0-10-1.734-10-3.5V17zm0-7v4m20-4v4"/>`),
		g.Group(children),
	)
}