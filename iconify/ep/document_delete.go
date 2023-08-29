package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentDelete(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M805.504 320L640 154.496V320h165.504zM832 384H576V128H192v768h640V384zM160 64h480l256 256v608a32 32 0 0 1-32 32H160a32 32 0 0 1-32-32V96a32 32 0 0 1 32-32zm308.992 546.304l-90.496-90.624l45.248-45.248l90.56 90.496l90.496-90.432l45.248 45.248l-90.496 90.56l90.496 90.496l-45.248 45.248l-90.496-90.496l-90.56 90.496l-45.248-45.248l90.496-90.496z"/>`),
		g.Group(children),
	)
}