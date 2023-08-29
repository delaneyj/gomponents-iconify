package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M209.5 0H0v209.5L81.5 128l104.7 104.7l46.5-46.5L128 81.5L209.5 0zm-23.3 279.3L81.5 384L0 302.5V512h209.5L128 430.5l104.7-104.7l-46.5-46.5zM302.5 0L384 81.5L279.3 186.2l46.5 46.5L430.5 128l81.5 81.5V0H302.5zm23.3 279.3l-46.5 46.5L384 430.5L302.5 512H512V302.5L430.5 384L325.8 279.3z"/>`),
		g.Group(children),
	)
}