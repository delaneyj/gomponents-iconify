package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.405 2.799c-.112-.44-.656-.799-1.21-.799H2.805c-.555 0-1.099.359-1.21.799L1.394 4h17.211l-.2-1.201zM19.412 5H.587a.58.58 0 0 0-.577.635l.923 11.669a.77.77 0 0 0 .766.696H18.3a.77.77 0 0 0 .766-.696l.923-11.669A.58.58 0 0 0 19.412 5zM8 14V9l4.383 2.5L8 14z"/>`),
		g.Group(children),
	)
}