package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrenchThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m27.84 6.775l-3.198 3.24l-2.872-.77L21 6.373l3.187-3.23a5.727 5.727 0 0 0-5.848 1.39a5.738 5.738 0 0 0-1.28 6.172l-9.64 9.64a3.971 3.971 0 0 0-.62-.062a3.932 3.932 0 0 0-3.934 3.933a3.932 3.932 0 1 0 7.865 0c0-.24-.03-.473-.07-.7l9.59-9.59a5.75 5.75 0 0 0 6.203-1.272a5.733 5.733 0 0 0 1.384-5.878zM6.8 25.145a.934.934 0 0 1-.936-.932a.932.932 0 1 1 .935.933z"/>`),
		g.Group(children),
	)
}