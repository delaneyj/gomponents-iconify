package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Deepnote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.22 11.506l.111.018c-.038-.006-.075-.011-.111-.018zm4.475 8.073c.636-2.816-2.172-4.8-6.955-4.814L.713 24h9.586c.132-.025.256-.056.384-.085c2.258-1.057 3.598-2.501 4.012-4.336zM10.299 24h.203l.021-.01c-.075.003-.148.008-.224.01zM24 11.319C24 3.15 18.711-.597 8.134.077L0 11.319h7.568c3.323 0 8.457.719 8.457 6.153c0 3.622-1.909 5.798-5.727 6.528c.099-.003.194-.009.291-.013l-.011.001l-.076.012h.912l.247-.077C19.885 23.27 24 19.07 24 11.319z"/>`),
		g.Group(children),
	)
}