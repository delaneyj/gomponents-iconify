package iwwa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortBy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 40 40"),
		g.Raw(`<path fill="currentColor" d="m20.062 24.84l-7.938 9.58V4.524a.75.75 0 0 0-1.5 0v29.854l-7.938-9.539a.75.75 0 1 0-1.154.959l9.285 11.157a.75.75 0 0 0 .577.271h.001c.223 0 .434-.1.577-.271l9.246-11.157a.75.75 0 0 0-1.156-.958zm18.405-10.638L29.183 3.045a.752.752 0 0 0-.576-.271h-.001c-.223 0-.435.1-.576.271l-9.247 11.157a.75.75 0 0 0 1.155.957l7.938-9.579v29.896a.75.75 0 0 0 1.5 0V5.621l7.938 9.539a.745.745 0 0 0 1.055.098a.75.75 0 0 0 .098-1.056z"/>`),
		g.Group(children),
	)
}