package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M330 0c197 0 337 145 337 357c0 197-140 360-337 360S0 554 0 357C0 149 133 0 330 0zm1 636c104 0 126-120 126-274c0-185-28-283-126-283c-113 0-122 132-122 297c0 154 29 260 122 260z"/>`),
		g.Group(children),
	)
}