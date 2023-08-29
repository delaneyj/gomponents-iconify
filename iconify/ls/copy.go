package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m572 22l105 109c10 11 19 34 19 49v323c0 15-12 29-28 29H481v-49h157c5 0 9-4 9-9V202c0-5-4-9-9-9H522c-5 0-10-4-10-9V60c0-5-4-9-9-9H313c-5 0-10 4-10 9v69c-10-4-23-6-33-6h-15V30c0-16 13-28 28-28h242c15 0 37 9 47 20zm-2 122h42c5 0 6-3 3-7l-48-49c-3-4-6-3-6 2v45c0 5 4 9 9 9zm-252 38l104 110c11 11 19 33 19 48v323c0 16-12 29-27 29H29c-16 0-29-13-29-29V190c0-15 13-29 29-29h241c15 0 37 10 48 21zM58 643h325c5 0 10-4 10-9V362c0-5-5-9-10-9H267c-5 0-9-4-9-9V220c0-5-4-9-9-9H58c-5 0-9 4-9 9v414c0 5 4 9 9 9zm302-345l-47-51c-4-4-7-1-7 4v44c0 5 5 9 10 9h41c5 0 7-3 3-6z"/>`),
		g.Group(children),
	)
}