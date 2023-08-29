package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M475 448V320L411 43Q401 0 368 0H112Q79 0 69 43L5 320v128q0 27 18.5 45.5T69 512h342q27 0 45.5-18.5T475 448zM103 60q5-17 22-17h230q17 0 22 17l51 241q-7-2-17-2H69q-10 0-17 2zm329 388q0 21-21 21H69q-21 0-21-21v-85q0-10 6-16t15-6h342q9 0 15 6t6 16v85zm-43-43q0 9-6 15.5t-15 6.5t-15-6.5t-6-15.5q0-8 6-14.5t15-6.5t15 6.5t6 14.5zm-85 0q0 9-6.5 15.5T283 427q-9 0-15.5-6.5T261 405q0-8 6.5-14.5T283 384q8 0 14.5 6.5T304 405z"/>`),
		g.Group(children),
	)
}