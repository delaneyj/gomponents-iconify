package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.176 2.9a.95.95 0 0 0-1.9 0v25.812a.95.95 0 1 0 1.9 0V18.107c.405.08.76.185 1.12.291c.789.233 1.596.47 2.988.47c1.429 0 2.944-.484 4.447-.965c1.477-.472 2.943-.94 4.3-.94c1.858 0 3.876.684 5.017 1.197c.241.108.525-.064.525-.329V5.494a.736.736 0 0 0-.425-.67c-4.297-1.97-6.75-1.169-9.283-.34c-.472.153-.947.309-1.437.447c-1.992.563-3.766.828-5.977.188a8.2 8.2 0 0 0-1.275-.36V2.9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}