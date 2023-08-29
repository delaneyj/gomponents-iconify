package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Logstash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#FEC514" d="M122.435 189.217H0V0h11.13c61.474 0 111.305 49.83 111.305 111.304v77.913Z"/><path fill="#343741" d="M116.87 306.086h5.565V189.217H0c0 64.545 52.324 116.87 116.87 116.87"/><path fill="#00BFB3" d="M150.261 306.086H256V189.217H150.261z"/>`),
		g.Group(children),
	)
}