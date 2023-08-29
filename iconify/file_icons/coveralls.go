package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coveralls(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 0v512h512V0zm252.625 40.75l51.797 150.594h159.234l-128.64 94.875l52.515 157.453l-136.25-98.766l-135.047 98.39l55.782-157.562l-134.641-94.39h160.062z"/>`),
		g.Group(children),
	)
}