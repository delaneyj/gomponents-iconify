package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Circleb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-68.5T68.5 769T0 512t68.5-257T255 68.5T512 0t257 68.5T955.5 255t68.5 257t-68.5 257T769 955.5T512 1024zm192-448q0-57-44-96q44-38 44-96q0-53-37.5-90.5T576 256H352q-13 0-22.5 9.5T320 288v448q0 13 9.5 22.5T352 768h224q53 0 90.5-37.5T704 640v-64zM576 704H400q-7 0-11.5-4.5T384 688V528q0-6 4.5-11t11.5-5h176q26 0 45 19t19 45v64q0 27-19 45.5T576 704zm0-256H400q-7 0-11.5-4.5T384 432v-96q0-6 4.5-11t11.5-5h176q26 0 45 19t19 45.5t-19 45t-45 18.5z"/>`),
		g.Group(children),
	)
}