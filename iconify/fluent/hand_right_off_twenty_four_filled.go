package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandRightOffTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 6.06L2.22 3.28a.75.75 0 1 1 1.06-1.06l18.5 18.5a.75.75 0 0 1-1.06 1.06l-4.306-4.305l-.107.124a11.25 11.25 0 0 0-1.562 2.338l-.133.267a3.25 3.25 0 0 1-2.907 1.797H9.537c-1.043 0-2.067-.504-2.623-1.459C6.168 19.264 5 16.934 5 14.755V6.06Zm3 3l-1-1v2.938a.5.5 0 0 0 1 0V9.061Zm2-2.242l-2-2V4.25a1 1 0 0 1 2 0v2.568Zm3 3l-2-2V2.997a1 1 0 1 1 2 0v6.821Zm5.384 5.384l-4.451-4.451A.497.497 0 0 0 14 10.5V4.25a1 1 0 1 1 2 0v7.768c.383-.21.844-.405 1.357-.504c.68-.132 1.453-.21 2.09-.134c.311.037.665.119.961.31c.333.215.592.575.592 1.06a.75.75 0 0 1-.29.592l-2.191 1.705l-.135.155Z"/>`),
		g.Group(children),
	)
}