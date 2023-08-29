package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkAsProtected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1193 1611L489 907l283-283l704 704l-283 283zm102-283L772 805L670 907l523 523l102-102zm369 592v-640h128v768H128V0h922L922 128H256v1792h1408zm305-1481l-215 87l245 245l-268 268l-72-73l-89 86l-582-582l89-86l-70-69l268-268l241 242l91-212l74-74l362 362l-74 74zm-275-242l-98 168l79 82l172-97l-153-153zm124 574l-543-543l-110 121l543 543l110-121z"/>`),
		g.Group(children),
	)
}