package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Belt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10.5651 30.3641L4.20117 36.728L11.2722 43.7991L17.6362 37.4351"/><path d="M21.8789 19.0503L36.7281 4.20105L43.7992 11.2721L28.95 26.1214"/><path d="M9.85855 29.6569C9.0775 28.8758 9.0775 27.6095 9.85855 26.8285L18.3438 18.3432C19.1249 17.5621 20.3912 17.5621 21.1723 18.3432L29.6575 26.8285C30.4386 27.6095 30.4386 28.8758 29.6575 29.6569L21.1723 38.1422C20.3912 38.9232 19.1249 38.9232 18.3438 38.1422L9.85855 29.6569Z"/><path d="M25.4148 22.5857L18.3438 29.6567"/><path d="M31.7785 9.15061L38.8496 16.2217"/><path d="M26.8283 14.1008L33.8994 21.1719"/><path d="M15.5154 21.1716L26.8291 32.4854"/></g>`),
		g.Group(children),
	)
}