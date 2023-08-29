package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MdEasel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M128 176h256v128H128z" fill="currentColor"/><path d="M448 96H64c-8.8 0-16 7.2-16 16v256c0 8.8 6.9 16 15.8 16H448c8.8 0 16-7.2 16-16V112c0-8.8-7.2-16-16-16zm-32 240H96V144h320v192z" fill="currentColor"/><path d="M80 464h57l22.5-64h-56.4z" fill="currentColor"/><path d="M279.4 48h-46.8l-11.5 32h69.8z" fill="currentColor"/><path d="M375 464h57l-23-64h-56.5z" fill="currentColor"/><path d="M232 400h48v32h-48z" fill="currentColor"/>`),
		g.Group(children),
	)
}