package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpinnerSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M6 2a2 2 0 1 1 3.999-.001A2 2 0 0 1 6 2zm6.359 6a1.641 1.641 0 0 1 3.282 0a1.641 1.641 0 0 1-3.282 0zm-1.602 4.243a1.486 1.486 0 1 1 2.971 0a1.486 1.486 0 0 1-2.971 0zM6.654 14a1.346 1.346 0 1 1 2.693.001A1.346 1.346 0 0 1 6.654 14zm-4.116-1.757a1.22 1.22 0 1 1 2.439 0a1.22 1.22 0 0 1-2.439 0zM.896 8a1.104 1.104 0 1 1 2.207 0A1.104 1.104 0 0 1 .896 8zm1.861-4.243a1 1 0 0 1 2 0a1 1 0 0 1-2 0zm11.297 0a1.811 1.811 0 1 1-3.623 0a1.811 1.811 0 0 1 3.623 0z"/>`),
		g.Group(children),
	)
}