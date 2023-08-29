package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Safety(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.842 3.865v7.434c0 4.047 3.205 7.327 7.158 7.327c3.953 0 7.158-3.28 7.158-7.327v-7.43L9.957 1.446L2.842 3.865ZM9.955 0L18.5 2.874v8.425C18.5 16.104 14.694 20 10 20s-8.5-3.896-8.5-8.701V2.874L9.955 0Zm4.419 6.485a.66.66 0 0 0-.948.057L9.11 11.511L6.322 8.426a.66.66 0 0 0-.948-.038a.698.698 0 0 0-.037.971l3.29 3.64a.66.66 0 0 0 .995-.01l4.807-5.534a.698.698 0 0 0-.055-.97Z"/>`),
		g.Group(children),
	)
}