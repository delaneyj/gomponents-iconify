package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Box(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="#FBF063" fill-rule="evenodd" d="M7 22L50 0l43 22l-43 21.001L7 22z" clip-rule="evenodd"/><path fill="#F29C1F" fill-rule="evenodd" d="M50.003 42.997L7 22v54.28l43.006 21.714l-.003-54.997z" clip-rule="evenodd"/><path fill="#F0C419" fill-rule="evenodd" d="M50 97.994L93.006 76.28V22L50.003 42.997L50 97.994z" clip-rule="evenodd"/><path fill="#F29C1F" fill-rule="evenodd" d="m27.036 11.705l42.995 21.498l2.263-1.105l-43.047-21.524z" clip-rule="evenodd" opacity=".5"/><path fill="#fff" fill-rule="evenodd" d="M21.318 14.674L63.3 36.505l15.99-7.809L35.788 7.271z" clip-rule="evenodd" opacity=".5"/><path fill="#fff" fill-rule="evenodd" d="m63.312 36.505l15.978-7.818v11l-15.978 8.817V36.505z" clip-rule="evenodd" opacity=".5"/>`),
		g.Group(children),
	)
}