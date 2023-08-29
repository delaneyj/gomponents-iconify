package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 768v768q0 1 9 27t22 67t30 89t30 90t24 73t13 38H0q2-7 12-37t25-73t30-91t29-88t23-67t9-28V768H0V640l960-480l960 480v128h-128zM286 640h1348L960 303L286 640zm1250 128v768h128V768h-128zm-256 0v768h128V768h-128zm-256 0v768h128V768h-128zm-256 0v768h128V768H768zm-256 0v768h128V768H512zm-256 768h128V768H256v768zm1486 256l-42-128H220l-42 128h1564z"/>`),
		g.Group(children),
	)
}