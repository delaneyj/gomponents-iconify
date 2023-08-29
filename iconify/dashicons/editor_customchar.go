package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditorCustomchar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 5.4c1.27 0 2.24.36 2.91 1.08c.66.71 1 1.76 1 3.13c0 1.28-.23 2.37-.69 3.27c-.47.89-1.27 1.52-2.22 2.12v2h6v-2h-3.69c.92-.64 1.62-1.34 2.12-2.34c.49-1.01.74-2.13.74-3.35c0-1.78-.55-3.19-1.65-4.22S11.92 3.54 10 3.54s-3.43.53-4.52 1.57c-1.1 1.04-1.65 2.44-1.65 4.2c0 1.21.24 2.31.73 3.33c.48 1.01 1.19 1.71 2.1 2.36H3v2h6v-2c-.98-.64-1.8-1.28-2.24-2.17c-.45-.89-.67-1.96-.67-3.22c0-1.37.33-2.41 1-3.13C7.75 5.76 8.72 5.4 10 5.4z"/>`),
		g.Group(children),
	)
}