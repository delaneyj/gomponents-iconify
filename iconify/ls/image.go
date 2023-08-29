package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Image(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 640V54h717v586H0zm644-194V127H73v271l156-93l204 111l83-49zM510 298c-32 0-57-26-57-58s25-58 57-58s57 26 57 58s-25 58-57 58z"/>`),
		g.Group(children),
	)
}