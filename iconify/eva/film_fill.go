package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilmFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaFilmFill0"><g id="evaFilmFill1"><path id="evaFilmFill2" fill="currentColor" d="M18.26 3H5.74A2.74 2.74 0 0 0 3 5.74v12.52A2.74 2.74 0 0 0 5.74 21h12.52A2.74 2.74 0 0 0 21 18.26V5.74A2.74 2.74 0 0 0 18.26 3ZM7 11H5V9h2Zm-2 2h2v2H5Zm14-2h-2V9h2Zm-2 2h2v2h-2Zm2-7.26V7h-2V5h1.26a.74.74 0 0 1 .74.74ZM5.74 5H7v2H5V5.74A.74.74 0 0 1 5.74 5ZM5 18.26V17h2v2H5.74a.74.74 0 0 1-.74-.74Zm14 0a.74.74 0 0 1-.74.74H17v-2h2Z"/></g></g>`),
		g.Group(children),
	)
}