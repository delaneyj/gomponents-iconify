package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Amd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 512L369.628 369.736V142.264H142.085L0 0h512v512zM142.085 369.736V165.004L0 307.226V512h204.651l142.22-142.264H142.084z"/>`),
		g.Group(children),
	)
}