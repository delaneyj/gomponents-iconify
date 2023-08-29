package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AndroidLogoLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M174 156a10 10 0 1 1-10-10a10 10 0 0 1 10 10Zm-82-10a10 10 0 1 0 10 10a10 10 0 0 0-10-10Zm146 22v24a14 14 0 0 1-14 14H32a14 14 0 0 1-14-14v-22.87a111.44 111.44 0 0 1 36.28-82.37L27.76 60.24a6 6 0 1 1 8.48-8.48l27.37 27.36a108.59 108.59 0 0 1 64-21.12h.39a109.17 109.17 0 0 1 64.58 20.93l27.18-27.17a6 6 0 0 1 8.48 8.48l-26.31 26.31c1.26 1.15 2.5 2.32 3.72 3.53A109.29 109.29 0 0 1 238 168Zm-12 0a98 98 0 0 0-98-98h-.35C73.81 70.19 30 114.66 30 169.13V192a2 2 0 0 0 2 2h192a2 2 0 0 0 2-2Z"/>`),
		g.Group(children),
	)
}