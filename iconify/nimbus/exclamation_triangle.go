package nimbus

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExclamationTriangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15.35 12.81L9 2.08a1.22 1.22 0 0 0-2 0L.65 12.81a1.14 1.14 0 0 0 1 1.69h12.66a1.14 1.14 0 0 0 1.04-1.69zm-13.66.55L8 2.64l6.31 10.72z"/><path fill="currentColor" d="M7.32 5.45h1.25V10H7.32z"/><ellipse cx="7.95" cy="11.9" fill="currentColor" rx=".67" ry=".7"/>`),
		g.Group(children),
	)
}