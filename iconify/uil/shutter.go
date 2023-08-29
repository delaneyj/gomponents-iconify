package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shutter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.07 4.93A10 10 0 1 0 4.93 19.07A10 10 0 1 0 19.07 4.93ZM18.23 7h-5.47l2.35-2.35A8.14 8.14 0 0 1 18.23 7ZM9 4.6a8.15 8.15 0 0 1 3.87-.54L9 7.93ZM7 5.77v5.47L5.19 9.43l-.54-.54A8.14 8.14 0 0 1 7 5.77ZM4.6 15a8.12 8.12 0 0 1-.54-3.87L7.93 15Zm1.17 2h5.47l-2.35 2.35A8.14 8.14 0 0 1 5.77 17ZM15 19.4a8.13 8.13 0 0 1-3.87.54L15 16.07Zm0-6.16L13.24 15h-2.49L9 13.24v-2.48L10.76 9h2.48L15 10.76Zm2 5v-5.48l2.35 2.35A8.14 8.14 0 0 1 17 18.23ZM16.07 9h3.33a8.13 8.13 0 0 1 .54 3.87Z"/>`),
		g.Group(children),
	)
}