package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmojiNeutral(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 .4A9.6 9.6 0 0 0 .4 10a9.6 9.6 0 1 0 19.2-.001C19.6 4.698 15.301.4 10 .4zm0 17.199a7.6 7.6 0 1 1 0-15.2a7.6 7.6 0 0 1 0 15.2zm2.5-11.348c-.828 0-1.5.783-1.5 1.749s.672 1.75 1.5 1.75c.829 0 1.5-.783 1.5-1.75s-.671-1.749-1.5-1.749zM7.501 9.75C8.329 9.75 9 8.967 9 8s-.672-1.75-1.5-1.75S6 7.033 6 8s.672 1.75 1.501 1.75zM13 12.25H7a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 0-1.5z"/>`),
		g.Group(children),
	)
}