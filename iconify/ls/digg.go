package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Digg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M132 410H90V297h42v113zm0-191H0v270h222V104h-90v115zm131-37h90v-78h-90v78zm0 307h90V219h-90v270zm264-79h-42V297h42v113zM395 525v79h222V219H395v270h132v36H395zm395-115h-42V297h42v113zM658 525v79h222V219H658v270h132v36H658z"/>`),
		g.Group(children),
	)
}