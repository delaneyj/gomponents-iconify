package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snews(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M842.5 532Q790 562 735 571.5t-95-2.5q10-42 45.5-83t88.5-71q193-110 250-93q-10 41-69.5 110.5t-112 99.5zM583 384q-12-31-4.5-78.5T613 207Q710 23 763 0q12 32-.5 111T723 240.5T658.5 329T583 384zM256 640q106 0 181 56t75 135.5t-75 136t-181 56.5t-181-56T0 832h128q0 26 37.5 45t90.5 19t90.5-19t37.5-45.5t-37.5-45T256 768q-106 0-181-56.5t-75-136T75 440t181-56t181 56t75 136H384q0-27-37.5-45.5T256 512t-90.5 18.5t-37.5 45t37.5 45.5t90.5 19zm135-400q-12-52-3.5-130T415 0q44 23 89 207q12 52 7 98t-25 79q-31-17-56.5-55T391 240z"/>`),
		g.Group(children),
	)
}