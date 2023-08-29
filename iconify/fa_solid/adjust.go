package fa_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adjust(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M8 256c0 136.966 111.033 248 248 248s248-111.034 248-248S392.966 8 256 8S8 119.033 8 256zm248 184V72c101.705 0 184 82.311 184 184c0 101.705-82.311 184-184 184z"/>`),
		g.Group(children),
	)
}