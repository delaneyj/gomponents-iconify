package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Padlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 7c0-1.7 1.3-3 3-3s3 1.3 3 3v2h2V7c0-2.8-2.2-5-5-5S7 4.2 7 7v2h2V7zm4.5 7.5c0-.8-.7-1.5-1.5-1.5s-1.5.7-1.5 1.5c0 .4.2.8.5 1.1V17c0 .6.4 1 1 1s1-.4 1-1v-1.4c.3-.3.5-.7.5-1.1z" opacity=".5"/><path fill="currentColor" d="M17 9H7c-1.7 0-3 1.3-3 3v7c0 1.7 1.3 3 3 3h10c1.7 0 3-1.3 3-3v-7c0-1.7-1.3-3-3-3zm-4 6.6V17c0 .6-.4 1-1 1s-1-.4-1-1v-1.4c-.3-.3-.5-.7-.5-1.1c0-.8.7-1.5 1.5-1.5s1.5.7 1.5 1.5c0 .4-.2.8-.5 1.1z"/>`),
		g.Group(children),
	)
}