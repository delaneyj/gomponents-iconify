package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpLeftArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m3 3l18 18M3 3c.676.676 1.923 1.11 3.039 1.379c1.49.359 3.036.424 4.559.252c1.182-.134 2.484-.4 3.009-.924M3 3c.676.676 1.11 1.923 1.379 3.039c.359 1.49.424 3.036.252 4.559c-.134 1.182-.4 2.484-.924 3.009"/>`),
		g.Group(children),
	)
}