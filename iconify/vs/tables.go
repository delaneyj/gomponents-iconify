package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tables(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M128 0Q76 0 38 38T0 128v512q0 52 38 90t90 38h512q52 0 90-38t38-90V128q0-52-38-90T640 0H128zm0 128h512v512H128V128zM1024 0q-52 0-90 38t-38 90v512q0 52 38 90t90 38h512q52 0 90-38t38-90V128q0-52-38-90t-90-38h-512zM128 896q-52 0-90 38t-38 90v512q0 52 38 90t90 38h512q52 0 90-38t38-90v-512q0-52-38-90t-90-38H128zm0 128h512v512H128v-512zm896-128q-52 0-90 38t-38 90v512q0 52 38 90t90 38h512q52 0 90-38t38-90v-512q0-52-38-90t-90-38h-512zm0 128h512v512h-512v-512z"/>`),
		g.Group(children),
	)
}