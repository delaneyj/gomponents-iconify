package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Brush(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M862.965 1020q-59 16-142.5-13.5t-145.5-87.5t-62-115q0-34 5-66q42-1 84.5-19t70.5-46q47-48 61-106q62 2 99 45q20 23 31.5 56.5t14.5 64t10.5 62t19 54t41 37t75.5 14.5q6 23-48 63.5t-114 56.5zm-466-392q-27-28-96.5-116t-145-193.5T32.965 121t-25-114t113.5 25.5t196 123.5t192.5 145.5t115 97.5q47 47 47 114t-47 114.5t-114 47.5t-114-47zm147-254q-2 81-46.5 126t-124.5 46q4 6 44.5 49t44.5 48q84 1 131.5-47t46.5-132q-5-4-47.5-44.5t-48.5-45.5z"/>`),
		g.Group(children),
	)
}