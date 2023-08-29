package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MutedSpeaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ff5a79"/><circle cx="32" cy="32" r="24" fill="#333"/><path fill="#fff" d="M41.5 28.1V13.9c0-.5-.4-.9-.9-.9s-.9.4-.9.9v.5l-14.1 9.9h-6.2c-1.2 0-2.3.9-2.3 2.1v11.3c0 1.2 1 2.1 2.3 2.1h6.2l14.1 9.9v.5c0 .5.4.9.9.9s.9-.4.9-.9V36c2.3 0 4.2-1.7 4.2-3.9s-1.9-4-4.2-4"/><path fill="#ff5a79" d="m9.23 13.474l4.243-4.242l41.294 41.294l-4.242 4.243z"/>`),
		g.Group(children),
	)
}