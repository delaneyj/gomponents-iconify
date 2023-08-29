package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobiledataOffRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 17.15l.85-.85q.3-.3.7-.288t.7.288q.3.3.312.713t-.287.712L9.7 20.3q-.15.15-.325.213T9 20.575q-.2 0-.375-.063T8.3 20.3l-2.575-2.575q-.3-.3-.287-.713t.312-.712q.3-.275.7-.288t.7.288l.85.85V10.8L2.1 4.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275L10 12.8v4.35Zm6-4l-2-2V6.8l-.9.9q-.275.275-.7.275t-.7-.275q-.275-.275-.275-.7t.275-.7l2.6-2.6q.15-.15.325-.212T15 3.425q.2 0 .375.063t.325.212l2.6 2.6q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-.9-.9v6.35Z"/>`),
		g.Group(children),
	)
}