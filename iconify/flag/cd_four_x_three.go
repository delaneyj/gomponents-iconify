package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CdFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#007fff" d="M0 0h640v480H0z"/><path fill="#f7d618" d="M28.8 96H96l20.8-67.2L137.6 96h67.2l-54.4 41.6l20.8 67.2l-54.4-41.6l-54.4 41.6l20.8-67.2L28.8 96zM600 0L0 360v120h40l600-360V0h-40"/><path fill="#ce1021" d="M640 0L0 384v96L640 96V0"/>`),
		g.Group(children),
	)
}