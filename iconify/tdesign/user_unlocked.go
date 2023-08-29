package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserUnlocked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0ZM18 14c-.69 0-1.25.56-1.25 1.25V16h5.748v6.5h-9V16h1.252v-.75a3.25 3.25 0 0 1 5.541-2.305l.71.705l-1.41 1.418l-.71-.705A1.243 1.243 0 0 0 18 14Zm-2.502 4v2.5h5V18h-5ZM8 16a4 4 0 0 0-4 4h7.55v2H2v-2a6 6 0 0 1 6-6h3.5v2H8Z"/>`),
		g.Group(children),
	)
}