package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseFreeOfChargeButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<rect width="48.984" height="48.984" x="11.591" y="11.617" fill="#d0cfce" rx="1.699"/><g fill="none" stroke="#000" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M30.593 17.728s-2.769 7.54-8.462 10.309"/><path stroke-linecap="round" stroke-linejoin="round" d="M50.518 23.498H27.355v18.54m-5.224.384h28.387M22.131 32.96h28.387"/><path stroke-linejoin="round" d="M33.595 23.073v19.349m6.23-19.349v19.349m6.154-19.349v19.349"/><path stroke-linecap="round" stroke-miterlimit="10" d="m26.043 46.926l-2.304 3.838m9.296-3.09l-.398 4.459m8.003-4.462l.313 4.465m6.952-5.012l2.271 3.858"/><rect width="48.984" height="48.984" x="11.591" y="11.617" stroke-miterlimit="10" rx="1.699"/></g>`),
		g.Group(children),
	)
}