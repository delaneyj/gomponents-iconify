package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><circle cx="16" cy="16" r="16" fill="#000"/><path fill="#FFF" d="M12.478 6.435v19.478H9V6h3.478v.435zM14.957 6h1.956c4.044 0 6.74 2.565 6.74 6.217c0 2.957-1.61 5.218-4.523 5.957l5.087 7.739H20.13L13.174 15.13h3.348c2.217 0 3.608-1 3.608-2.956c0-1.957-1.39-2.957-3.608-2.957H15V6h-.043z"/></g>`),
		g.Group(children),
	)
}