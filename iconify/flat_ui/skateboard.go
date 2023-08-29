package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skateboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="#E64B3C" d="M50 0c27.613 0 50 22.386 50 50s-22.387 50-50 50C22.386 100 0 77.614 0 50S22.386 0 50 0z"/><defs><circle id="flatUiSkateboard0" cx="50" cy="50" r="50"/></defs><clipPath id="flatUiSkateboard1"><use href="#flatUiSkateboard0"/></clipPath><g clip-path="url(#flatUiSkateboard1)"><path fill="#C03A2B" d="m35 25l-67 67l27 27h57V29l-17-4z"/><path fill="#EBECED" d="M50 20c14.359 0 26 11.641 26 26v43c0 14.359-11.641 26-26 26s-26-11.641-26-26V46c0-14.359 11.641-26 26-26z"/><path fill="#A1A5A9" d="M46 43h8a3 3 0 0 1 3 3v13a3 3 0 0 1-3 3h-8a3 3 0 0 1-3-3V46a3 3 0 0 1 3-3z"/><path fill="#2D3E50" d="M50 45a1 1 0 1 1 .002 1.998A1 1 0 0 1 50 45z"/><path fill="#BDC2C7" d="M51 49H36v3h10v4a3 3 0 0 0 3 3h2a3 3 0 0 0 3-3v-4h10v-3H51z"/><path fill="#2D3E50" d="M50 53a2 2 0 1 1-.001 4.001A2 2 0 0 1 50 53z"/><path fill="#35495E" d="M66 43h4a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3h-4a3 3 0 0 1-3-3v-9a3 3 0 0 1 3-3z"/><path fill="#2D3E50" d="M66 43h4v15h-4V43z"/><path fill="#35495E" d="M30 43h4a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3h-4a3 3 0 0 1-3-3v-9a3 3 0 0 1 3-3z"/><path fill="#2D3E50" d="M30 43h4v15h-4V43z"/><path fill="#F29C1F" d="M50 69c-6.075 0-11 4.925-11 11c0 3.859 1.992 7.25 5 9.213V91.5a1.5 1.5 0 1 0 3 0a1.5 1.5 0 1 0 3 0a1.5 1.5 0 0 0 3 0a1.5 1.5 0 0 0 3 0v-2.287c3.008-1.963 5-5.354 5-9.213c0-6.075-4.925-11-11-11zm-3.5 14a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5zm7 0a2.5 2.5 0 1 1 0-5a2.5 2.5 0 0 1 0 5z"/><path fill="#919598" d="M37 53h9v-2h-9v2zm17-2v2h9v-2h-9z"/><path fill="#D4D5D6" d="M37 51h6v2h-6zm20 0h6v2h-6z"/><path fill="#AAAFB4" d="M54 51h9v1h-9zm-17 0h9v1h-9z"/></g>`),
		g.Group(children),
	)
}