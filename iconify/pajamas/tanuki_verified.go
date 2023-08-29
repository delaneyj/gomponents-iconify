package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TanukiVerified(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m15.712 6.337l.022.058a4.043 4.043 0 0 1-1.343 4.671l-.01.008l-3.336 2.497l-1.643 1.242l-.999.755a.675.675 0 0 1-.813 0l-.999-.755l-1.643-1.242l-3.313-2.483l-.018-.014l-.008-.007A4.043 4.043 0 0 1 .267 6.395l.023-.057L2.466.657a.565.565 0 0 1 .475-.358a.576.576 0 0 1 .613.416l1.47 4.5h5.953l1.47-4.5a.575.575 0 0 1 1.087-.058l2.178 5.68ZM11.28 8.53a.75.75 0 0 0-1.06-1.06L7.5 10.19L6.28 8.97a.75.75 0 1 0-1.06 1.06l1.75 1.75a.75.75 0 0 0 1.06 0l3.25-3.25Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}