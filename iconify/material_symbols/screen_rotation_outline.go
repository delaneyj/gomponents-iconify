package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenRotationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.4 19.45L4.55 11.6q-.575-.575-.575-1.35T4.55 8.9L8.9 4.55q.575-.575 1.35-.575t1.35.575l7.85 7.85q.575.575.575 1.35t-.575 1.35l-4.35 4.35q-.575.575-1.35.575t-1.35-.575ZM13.75 18L18 13.75L10.25 6L6 10.25L13.75 18ZM12 24q-2.475 0-4.663-.937t-3.825-2.575Q1.875 18.85.937 16.663T0 12h2q0 1.775.6 3.4t1.663 2.925q1.062 1.3 2.537 2.213t3.225 1.287L7.4 19.2l1.4-1.4l5.9 5.9q-.65.15-1.337.225T12 24Zm10-12q0-1.775-.6-3.4t-1.663-2.925q-1.062-1.3-2.537-2.212t-3.225-1.288L16.6 4.8l-1.4 1.4L9.3.3q.65-.15 1.338-.225T12 0q2.475 0 4.663.938t3.825 2.575q1.637 1.637 2.575 3.825T24 12h-2Zm-10 0Z"/>`),
		g.Group(children),
	)
}