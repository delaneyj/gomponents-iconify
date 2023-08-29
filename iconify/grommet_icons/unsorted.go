package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unsorted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m15.204 15.321l1.597-1.597l.707.707l-2.45 2.45l-.354.354l-.353-.353l-2.45-2.45l.707-.708l1.596 1.597V7.217h1v8.104Zm-5.9-6.407v8.104h1V8.914l1.597 1.597l.707-.707l-2.45-2.45L9.803 7l-.354.354L7 9.804l.707.707l1.597-1.597Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}