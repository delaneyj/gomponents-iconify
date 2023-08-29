package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nodejs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m321 448l128-64l128 64v128l-128 64l-128-64V448zm552 3L683 576q-20 0-31-13t-11-32V400L449 288L258 400v224l443 259l-192 127q-25 14-60.5 14t-60.5-14L26 771q-28-16-25-39V292q-3-23 25-38L388 14q25-14 60.5-14T509 14l363 240q25 14 24 34h1v124q4 23-24 39z"/>`),
		g.Group(children),
	)
}