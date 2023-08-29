package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TargetTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.998 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm-6-2a6 6 0 1 1 12 0a6 6 0 0 1-12 0Zm6-4.5a4.5 4.5 0 1 0 0 9a4.5 4.5 0 0 0 0-9ZM1.996 12c0-5.524 4.478-10.002 10.002-10.002C17.522 1.998 22 6.476 22 12c0 5.524-4.478 10.002-10.002 10.002c-5.524 0-10.002-4.478-10.002-10.002Zm10.002-8.502a8.502 8.502 0 1 0 0 17.004a8.502 8.502 0 0 0 0-17.004Z"/>`),
		g.Group(children),
	)
}