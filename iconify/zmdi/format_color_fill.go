package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatColorFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M353 191q10 9 10 22.5T353 236L236 353q-9 10-22.5 10T191 353L73 236q-9-9-9-22.5t9-22.5L183 81l-51-51l31-30zm-242 22h205L213 111zm294 32q43 47 43 75q0 18-12.5 30.5t-30 12.5t-30-12.5T363 320q0-13 10.5-31.5T395 258zM0 427h512v85H0v-85z"/>`),
		g.Group(children),
	)
}