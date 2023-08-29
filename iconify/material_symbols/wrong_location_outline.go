package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WrongLocationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-4.025-3.425-6.012-6.362T4 10.2q0-3.75 2.413-5.975T12 2q.25 0 .488.013t.512.062V4.1q-.25-.05-.5-.075T12 4Q9.475 4 7.737 5.737T6 10.2q0 1.775 1.475 4.063T12 19.35q3.05-2.8 4.525-5.088T18 10.2q0-.05-.013-.1t-.012-.1h2q0 .05.013.1t.012.1q0 2.5-1.988 5.438T12 22Zm0-11.25Zm4.875-2.7l2.1-2.1l2.1 2.1l1.4-1.4l-2.1-2.1l2.1-2.1l-1.4-1.4l-2.1 2.1l-2.1-2.1l-1.4 1.4l2.1 2.1l-2.1 2.1l1.4 1.4ZM12 12q.825 0 1.413-.588T14 10q0-.825-.588-1.413T12 8q-.825 0-1.413.588T10 10q0 .825.588 1.413T12 12Z"/>`),
		g.Group(children),
	)
}