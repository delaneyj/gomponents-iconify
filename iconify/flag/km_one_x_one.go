package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KmOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagKm1x10"><path fill-opacity=".7" d="M0 0h416.3v416.3H0z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagKm1x10)" transform="scale(1.23)"><path fill="#ff0" d="M0 0h625v104H0z"/><path fill="#fff" d="M0 104h625v104.1H0z"/><path fill="#be0027" d="M0 208.1h625v104H0z"/><path fill="#3b5aa3" d="M0 312.2h625v104H0z"/><path fill="#239e46" d="M0 0v416.2l310.4-207.5L0 0z"/><path fill="#fff" d="M127.8 115c-69.2-3.5-100.7 51.6-100.6 94.2c-.2 50.4 47.6 92 91.7 89.4A100 100 0 0 1 65.8 209a98.3 98.3 0 0 1 62-94z"/><path fill="#fff" d="m126.8 160.2l-9.8-7.6l-11.8 3.7l4.2-11.6l-7.1-10.1l12.3.4l7.4-10l3.4 12l11.8 3.9l-10.3 7zm-.2 42.3l-9.8-7.6l-11.8 3.7l4.2-11.6l-7.2-10.1l12.4.4l7.4-10l3.4 12l11.8 4l-10.3 6.9zm.2 42.8l-9.8-7.6l-11.8 3.7l4.2-11.7l-7.1-10l12.3.4l7.4-10l3.4 12l11.8 3.9l-10.3 6.9zm-.2 43.1l-9.8-7.6l-11.8 3.7l4.2-11.6l-7.2-10.1l12.4.4l7.4-10l3.4 12l11.8 4l-10.3 6.9z"/></g>`),
		g.Group(children),
	)
}