package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Source(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M334.65 484.494c27.29-25.657 49.663-57.373 64.799-94.23c1.043-2.517 105.582-257.15 105.85-257.83c24.405-62.001-20.837-118.76-87.297-116.975H162.851c-52.653-2.677-104.533 11.644-149.395 37.47c-33.224 21.245 1.986 67.57 31.167 40.33C277.276-22.119 481.432 265.707 299.676 448.171l.197.277c-34.711 26.13 8.513 69.963 34.777 36.045z"/>`),
		g.Group(children),
	)
}