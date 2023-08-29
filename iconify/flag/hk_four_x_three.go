package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HkFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#EC1B2E" d="M0 0h640v480H0"/><path id="flagHk4x30" fill="#fff" d="M346.3 103.1C267 98 230.6 201.9 305.6 240.3c-26-22.4-20.6-55.3-10.1-72.4l1.9 1.1c-13.8 23.5-11.2 52.7 11.1 71c-12.7-12.3-9.5-39 12.1-48.9s23.6-39.3 16.4-49.1q-14.7-25.6 9.3-38.9zM307.9 164l-4.7 7.4l-1.8-8.6l-8.6-2.3l7.8-4.3l-.6-8.9l6.5 6.1l8.3-3.3l-3.7 8.1l5.6 6.8z"/><use href="#flagHk4x30" transform="rotate(72 312.5 243.5)"/><use href="#flagHk4x30" transform="rotate(144 312.5 243.5)"/><use href="#flagHk4x30" transform="rotate(216 312.5 243.5)"/><use href="#flagHk4x30" transform="rotate(288 312.5 243.5)"/>`),
		g.Group(children),
	)
}