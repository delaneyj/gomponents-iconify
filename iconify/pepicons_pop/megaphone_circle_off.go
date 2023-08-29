package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MegaphoneCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="m10.4 13.054l5.6 1.527V8.62l-5.6 1.527v2.908Zm-2 .764a1 1 0 0 0 .737.965l7.6 2.072A1 1 0 0 0 18 15.891V7.309a1 1 0 0 0-1.263-.965l-7.6 2.073a1 1 0 0 0-.737.965v4.436Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8.016 12.8v-2.4H6.97a2.429 2.429 0 0 0 .004 2.4h1.043Zm1 2a1 1 0 0 0 1-1V9.4a1 1 0 0 0-1-1H6.253a.55.55 0 0 0-.4.172c-1.602 1.691-1.595 4.353-.002 6.053a.555.555 0 0 0 .405.175h2.76Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M9.424 15H8.84l-.766 3h.583l.767-3Zm-.584-2a2 2 0 0 0-1.938 1.506l-1.084 4.247A1 1 0 0 0 6.788 20h1.87a2 2 0 0 0 1.937-1.505l.767-3A2 2 0 0 0 9.424 13H8.84Zm13.192-5.555a1 1 0 0 1-.277 1.387l-1.5 1a1 1 0 0 1-1.11-1.664l1.5-1a1 1 0 0 1 1.387.277ZM18.7 11.6a1 1 0 0 1 1-1h1.5a1 1 0 0 1 0 2h-1.5a1 1 0 0 1-1-1Zm.234 1.909a1 1 0 0 1 1.409-.123l1.38 1.16a1 1 0 0 1-1.286 1.531l-1.38-1.16a1 1 0 0 1-.123-1.408Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}