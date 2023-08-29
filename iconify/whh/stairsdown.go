package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stairsdown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 512H672q-13 0-22.5-9.5T640 480v-64q0-13 9.5-22.5T672 384h135L521 100q-9-9-9-22.5t9-22.5l46-46q9-9 22.5-9T613 9l283 282V160q0-13 9.5-22.5T928 128h64q13 0 22.5 9.5t9.5 22.5v320q0 13-9.5 22.5T992 512zm-353 90q1 2 1 6v160h160q12 0 20.5 7.5T831 794q1 2 1 6v160h160q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5H800q-12 0-20.5-7.5T769 998q-1-2-1-6V832H608q-12 0-20.5-7.5T577 806q-1-2-1-6V640H416q-12 0-20.5-7.5T385 614q-1-2-1-6V448H224q-12 0-20.5-7.5T193 422q-1-2-1-6V256H32q-13 0-22.5-9.5T0 224t9.5-22.5T32 192h192q12 0 20.5 7.5T255 218q1 2 1 6v160h160q12 0 20.5 7.5T447 410q1 2 1 6v160h160q12 0 20.5 7.5T639 602z"/>`),
		g.Group(children),
	)
}