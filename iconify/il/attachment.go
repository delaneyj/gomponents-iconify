package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Attachment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M440 240q10 0 17 6t6 17v255q0 50-19 93t-54 75t-78 49t-96 14q-46-3-85-24t-69-53t-45-74t-17-88V186q0-79 34-125t87-52q31-4 60 5t50 29t35 46t12 58v348q0 10-6 16t-17 7h-47q-9 0-16-7t-7-16V150q0-17-10-32t-28-17q-22-3-38 11t-16 35v366q0 27 9 51t25 43t38 33t48 16q32 4 61-6t50-28t35-46t12-58V263q0-10 7-17t16-6h46z"/>`),
		g.Group(children),
	)
}