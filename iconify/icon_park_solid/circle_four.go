package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSCircleFour0"><g fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill-rule="evenodd" d="M24 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0 34a3 3 0 1 0 0-6a3 3 0 0 0 0 6ZM7 27a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm34 0a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z" clip-rule="evenodd"/><path d="M24.197 15.744a8.99 8.99 0 0 0 7.438-3.931a14.528 14.528 0 0 1 4.686 4.498A8.995 8.995 0 0 0 32 24a8.997 8.997 0 0 0 4.527 7.811a14.53 14.53 0 0 1-4.426 4.532a8.998 8.998 0 0 0-7.905-4.694c-3.4 0-6.36 1.885-7.89 4.668a14.528 14.528 0 0 1-4.494-4.683a8.99 8.99 0 0 0 3.93-7.438a8.99 8.99 0 0 0-3.907-7.422a14.526 14.526 0 0 1 4.94-4.938a8.99 8.99 0 0 0 7.42 3.907Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSCircleFour0)"/>`),
		g.Group(children),
	)
}