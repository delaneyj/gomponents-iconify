package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockwiseVerticalArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiClockwiseVerticalArrows0" d="m26.675 17l8.215 9.452h-6.024l.236 25.086s-.746 6.721 7.468 2.24l1.494 2.987c-13.442 7.468-13.442-5.974-13.442-5.974l-.137-24.34H18.46L26.675 17"/><path id="openmojiClockwiseVerticalArrows1" d="m44.785 59.025l-8.214-9.452h6.024l-.237-25.087s.747-6.72-7.468-2.24l-1.493-2.987c13.442-7.468 13.442 5.974 13.442 5.974l.137 24.34H53l-8.215 9.452"/></defs><g fill="#92d3f5"><use href="#openmojiClockwiseVerticalArrows0"/><use href="#openmojiClockwiseVerticalArrows1"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><use href="#openmojiClockwiseVerticalArrows0"/><use href="#openmojiClockwiseVerticalArrows1"/></g>`),
		g.Group(children),
	)
}