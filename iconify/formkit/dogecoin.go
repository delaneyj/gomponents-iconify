package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dogecoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.72 1.5H4V6H2l-.04 2H4v4.5h3.91c1.68 0 2.08-.2 2.89-.59c.06-.03.13-.06.2-.1c.95-.46 1.69-1.1 2.21-1.92c.53-.83.79-1.79.79-2.89s-.26-2.06-.79-2.89c-.53-.83-1.26-1.47-2.21-1.93c-.95-.46-2.04-.68-3.28-.68ZM6.75 8v2.41h.83c1.14 0 2.05-.31 2.73-.92c.68-.61 1.02-1.44 1.02-2.49s-.34-1.88-1.02-2.49c-.68-.61-1.59-.92-2.73-.92h-.83V6h2.58v2H6.75Z"/>`),
		g.Group(children),
	)
}