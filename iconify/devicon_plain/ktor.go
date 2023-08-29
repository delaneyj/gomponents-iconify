package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ktor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M47.543 14.629L14.629 47.543l32.414 32.414V47.043h32.914L47.543 14.629zm.5 33.414v31.914h31.914V48.043H48.043zm32.914 0v32.914H48.043l32.414 32.414l32.914-32.914l-32.414-32.414z"/>`),
		g.Group(children),
	)
}