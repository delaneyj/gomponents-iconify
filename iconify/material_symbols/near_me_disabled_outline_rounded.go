package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NearMeDisabledOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.675 20.425L10.05 13.95l-6.475-2.625q-.325-.125-.475-.387t-.15-.538q0-.275.163-.537t.487-.388l4.275-1.6L3.5 3.5q-.3-.3-.3-.713t.3-.712q.3-.3.713-.3t.712.3l17 17q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3l-4.375-4.35l-1.6 4.275q-.125.325-.388.488t-.537.162q-.275 0-.537-.15t-.388-.475Zm5-8.45L16.1 10.4l1.5-4l-4 1.5l-1.575-1.575L18.95 3.75q.3-.125.588-.05t.487.275q.2.2.275.487t-.05.588l-2.575 6.925ZM13.55 17.3l1.025-2.725l-5.15-5.15L6.7 10.45l4.9 1.95l1.95 4.9Zm1.3-8.15ZM12 12Z"/>`),
		g.Group(children),
	)
}