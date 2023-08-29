package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterN(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M20.2 17.5h6.4l11.5 20.2V17.5h5.6v29h-6.1L25.9 25.9v20.6h-5.6l-.1-29"/>`),
		g.Group(children),
	)
}