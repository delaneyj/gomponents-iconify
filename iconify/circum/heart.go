package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 20.043a.977.977 0 0 1-.7-.288L4.63 13.08a5.343 5.343 0 0 1 1.423-8.567A5.266 5.266 0 0 1 12 5.371a5.272 5.272 0 0 1 5.947-.858a5.343 5.343 0 0 1 1.423 8.567l-6.676 6.675a.977.977 0 0 1-.694.288ZM8.355 4.963a4.015 4.015 0 0 0-1.844.437a4.4 4.4 0 0 0-2.389 3.243a4.345 4.345 0 0 0 1.215 3.73l6.675 6.675l6.651-6.675a4.345 4.345 0 0 0 1.215-3.73A4.4 4.4 0 0 0 17.489 5.4a4.338 4.338 0 0 0-4.968.852a.744.744 0 0 1-1.042 0a4.474 4.474 0 0 0-3.124-1.289Z"/>`),
		g.Group(children),
	)
}