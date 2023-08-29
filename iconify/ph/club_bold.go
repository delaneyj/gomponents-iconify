package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClubBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M186.79 84.06a60 60 0 1 0-117.58 0a60 60 0 1 0 19 117.73l-4 19.86A12 12 0 0 0 96 236h64a12 12 0 0 0 11.77-14.35l-4-19.86a60 60 0 1 0 19-117.73ZM184 180a35.88 35.88 0 0 1-24.78-9.88a12 12 0 0 0-20 11.06l6.14 30.82h-34.72l6.16-30.82a12 12 0 0 0-20-11.06a36 36 0 1 1-12.65-60a12 12 0 0 0 14-18A35.68 35.68 0 0 1 92 72a36 36 0 1 1 65.88 20.08a12 12 0 0 0 14 18A36 36 0 1 1 184 180Z"/>`),
		g.Group(children),
	)
}