package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalfTwelveRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m6 9.428l-2.223 1.169a.8.8 0 0 1-1.16-.844l.424-2.476l-1.799-1.753a.8.8 0 0 1 .444-1.365l2.486-.36l1.111-2.253A.79.79 0 0 1 6 1.1a.79.79 0 0 1 .718.446L7.83 3.798l2.486.361a.8.8 0 0 1 .443 1.365L8.96 7.277l.425 2.476a.8.8 0 0 1-1.16.844L6 9.428Zm0-7.075v6.049a.8.8 0 0 1 .373.092l1.958 1.029l-.374-2.18a.8.8 0 0 1 .23-.708L9.771 5.09l-2.189-.318a.8.8 0 0 1-.602-.438L6 2.352v.001Z"/>`),
		g.Group(children),
	)
}