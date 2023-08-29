package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenOffTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.28 2.22a.75.75 0 1 0-1.06 1.06l1.263 1.264A3.235 3.235 0 0 0 3 6.25v11.5A3.25 3.25 0 0 0 6.25 21h11.5c.626 0 1.21-.177 1.706-.483l1.263 1.264a.75.75 0 0 0 1.061-1.061L3.28 2.22Zm14.647 16.768a1.254 1.254 0 0 1-.177.012H6.25C5.56 19 5 18.44 5 17.75V6.25c0-.06.004-.12.012-.177l12.915 12.915Zm-4.913-9.156l1.154 1.154a.995.995 0 0 0 .54-.279L19 6.414V10a1 1 0 1 0 2 0V4a1 1 0 0 0-1-1h-6a1 1 0 1 0 0 2h3.586l-4.293 4.293a.995.995 0 0 0-.279.54ZM21 17.818l-1.999-2V14a1 1 0 1 1 2 0v3.818ZM6.183 3l1.999 2H10a1 1 0 1 0 0-2H6.183Z"/>`),
		g.Group(children),
	)
}