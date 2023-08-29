package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MegaphoneCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopMegaphoneCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="m10.4 13.054l5.6 1.527V8.62l-5.6 1.527v2.908Zm-2 .764a1 1 0 0 0 .737.965l7.6 2.072A1 1 0 0 0 18 15.891V7.309a1 1 0 0 0-1.263-.965l-7.6 2.073a1 1 0 0 0-.737.965v4.436Z"/><path d="M8.016 12.8v-2.4H6.97a2.429 2.429 0 0 0 .004 2.4h1.043Zm1 2a1 1 0 0 0 1-1V9.4a1 1 0 0 0-1-1H6.253a.55.55 0 0 0-.4.172c-1.602 1.691-1.595 4.353-.002 6.053a.555.555 0 0 0 .405.175h2.76Z"/><path d="M9.424 15H8.84l-.766 3h.583l.767-3Zm-.584-2a2 2 0 0 0-1.938 1.506l-1.084 4.247A1 1 0 0 0 6.788 20h1.87a2 2 0 0 0 1.937-1.505l.767-3A2 2 0 0 0 9.424 13H8.84Zm13.192-5.555a1 1 0 0 1-.277 1.387l-1.5 1a1 1 0 0 1-1.11-1.664l1.5-1a1 1 0 0 1 1.387.277ZM18.7 11.6a1 1 0 0 1 1-1h1.5a1 1 0 0 1 0 2h-1.5a1 1 0 0 1-1-1Zm.234 1.909a1 1 0 0 1 1.409-.123l1.38 1.16a1 1 0 0 1-1.286 1.531l-1.38-1.16a1 1 0 0 1-.123-1.408Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopMegaphoneCircleFilled0)"/></g>`),
		g.Group(children),
	)
}