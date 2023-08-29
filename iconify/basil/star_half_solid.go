package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarHalfSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.5 2.462c0-.26-.208-.488-.45-.395a.987.987 0 0 0-.504.431L8.832 7.175a1.25 1.25 0 0 1-.894.61l-5.086.767c-.855.13-1.154 1.208-.489 1.76l3.79 3.138c.35.29.515.75.43 1.197l-.992 5.205a1 1 0 0 0 1.449 1.072l4.79-2.522c.046-.025.094-.046.143-.065c.271-.101.527-.322.527-.613V2.462Z"/>`),
		g.Group(children),
	)
}