package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoChat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 49.109v854.518h247.266v247.264L632.74 903.625H1200V49.109H0zm263.086 214.673h483.398v149.049l190.43-149.049v398.439l-190.43-149.049v149.047H263.086V263.782z"/>`),
		g.Group(children),
	)
}