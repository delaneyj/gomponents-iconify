package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DieOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M83.921 23.292a7.44 7.44 0 0 0-7.213-7.264v-.011H23.292v.008a7.443 7.443 0 0 0-7.267 7.267h-.009v53.416h.009a7.44 7.44 0 0 0 7.267 7.267v.008h53.416v-.009a7.44 7.44 0 0 0 7.267-7.267h.009V23.292h-.063zM50 57.621a7.621 7.621 0 1 1 0-15.24a7.621 7.621 0 0 1 0 15.24z"/>`),
		g.Group(children),
	)
}