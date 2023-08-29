package academicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrossrefSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M48 32C21.5 32 0 53.5 0 80v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V80c0-26.5-21.5-48-48-48H48zm287.379 64v95.803l-68.983 22.717l-142.271-48.877L335.379 96zM112.62 169.43l142.246 48.877l-142.246 48.877V169.43zm153.775 52.664l57.46 18.935l-142.272 48.858l-57.44-18.916l142.252-48.877zm68.983 22.722v97.754l-142.272-48.877l142.272-48.877zM181.584 297.48l142.271 48.877L112.621 416v-95.803l68.963-22.717z"/>`),
		g.Group(children),
	)
}