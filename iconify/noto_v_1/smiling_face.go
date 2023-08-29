package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmilingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#fcc21b" d="M63.79 8.64C1.48 8.64 0 78.5 0 92.33c0 13.83 28.56 25.03 63.79 25.03c35.24 0 63.79-11.21 63.79-25.03c0-13.83-1.47-83.69-63.79-83.69z"/><path fill="#2f2f2f" d="M42.21 66.47c-4.49.04-8.17-4.27-8.22-9.62c-.05-5.37 3.55-9.75 8.04-9.79c4.48-.04 8.17 4.27 8.22 9.64c.05 5.36-3.55 9.73-8.04 9.77zm44.11 0c4.48-.01 8.11-4.36 8.1-9.71c-.01-5.37-3.66-9.7-8.14-9.69c-4.49.01-8.13 4.36-8.12 9.73c.02 5.35 3.67 9.68 8.16 9.67z"/><path fill="none" stroke="#2f2f2f" stroke-linecap="round" stroke-miterlimit="10" stroke-width="6" d="M36.19 81.17c13.73 16.67 41.89 16.67 55.62 0"/>`),
		g.Group(children),
	)
}