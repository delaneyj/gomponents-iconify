package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EOneThousandFiveHundredFortySeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 23.48c2.379.872 8.613 3.914 8.613 3.914s9.184-6.324 13.46-2.74s-1.897 10.299-1.897 10.299c1.837.813 6.806 4.878 8.432 6.444c2.56-4.457 14.834-29.071 8.69-33.589S17.48 12.126 4.5 23.48Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.312 11.84c10.26 2.705 13.312 9.932 12.589 16.517c-.609 5.545 2.409 6.757 2.409 6.757"/>`),
		g.Group(children),
	)
}