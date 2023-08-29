package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatStatus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12a9.97 9.97 0 0 0 1.3 4.935l-1.249 3.749a1 1 0 0 0 1.265 1.265l3.749-1.25A9.96 9.96 0 0 0 12 22c5.523 0 10-4.477 10-10S17.523 2 12 2zm0 6c-.902 0-1.731.297-2.4.8a1 1 0 1 1-1.2-1.6a6 6 0 0 1 8.4 8.4a1 1 0 0 1-1.598-1.2A4 4 0 0 0 12 8zm-5 3a1 1 0 0 1 1 1a4 4 0 0 0 4 4a1 1 0 1 1 0 2a6 6 0 0 1-6-6a1 1 0 0 1 1-1zm5-1a2 2 0 1 0 0 4a2 2 0 0 0 0-4z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}