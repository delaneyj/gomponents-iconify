package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tooth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M640 704v51l-1 52l-4 55l-7 49.5l-11.5 45.5l-17 34l-24 25l-31.5 8q-8 0-13.5-3.5t-9-12t-5.5-16t-3-23.5t-1-26v-47q0-62-.5-91.5t-4-75T496 663t-22-45t-36.5-33t-53.5-9t-53.5 9t-36.5 33t-22 45t-11.5 66.5t-4 75t-.5 91.5v47l-1 26l-3 23.5l-5.5 16l-9 12l-13.5 3.5q-17 0-31.5-8t-24-25t-17-34t-11.5-45.5t-7-49.5t-4-55t-1-52v-51q0-42-20-107T64 472.5T20 335T0 192q0-80 56-136T192 0q32 0 96 32t96 32t96-32t96-32q79 0 135.5 56T768 192q0 65-20 143t-44 137.5T660 597t-20 107z"/>`),
		g.Group(children),
	)
}