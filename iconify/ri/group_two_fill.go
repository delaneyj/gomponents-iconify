package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 19.748V16.4c0-1.283.995-2.292 2.467-2.868A8.48 8.48 0 0 0 9.5 13c-1.89 0-3.636.617-5.047 1.66A8.017 8.017 0 0 0 10 19.748Zm8.88-3.662C18.485 15.553 17.17 15 15.5 15c-2.006 0-3.5.797-3.5 1.4V20a7.996 7.996 0 0 0 6.88-3.914ZM9.55 11.5a2.25 2.25 0 1 0 0-4.5a2.25 2.25 0 0 0 0 4.5Zm5.95 1a2 2 0 1 0 0-4a2 2 0 0 0 0 4ZM12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Z"/>`),
		g.Group(children),
	)
}