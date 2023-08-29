package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaxiDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M184 184h40v24a8 8 0 0 1-8 8h-24a8 8 0 0 1-8-8ZM32 208a8 8 0 0 0 8 8h24a8 8 0 0 0 8-8v-24H32ZM194.3 68a8 8 0 0 0-6.94-4H68.64a8 8 0 0 0-6.94 4L32 120h192Z" opacity=".2"/><path d="M240 112h-11.36l-27.39-47.94A16 16 0 0 0 187.36 56h-21.94l-12-29.94A15.93 15.93 0 0 0 138.58 16h-21.16a15.93 15.93 0 0 0-14.86 10.06L90.58 56H68.64a16 16 0 0 0-13.89 8.06L27.36 112H16a8 8 0 0 0 0 16h8v80a16 16 0 0 0 16 16h24a16 16 0 0 0 16-16v-16h96v16a16 16 0 0 0 16 16h24a16 16 0 0 0 16-16v-80h8a8 8 0 0 0 0-16ZM117.42 32h21.16l9.6 24h-40.36ZM68.64 72h118.72l22.85 40H45.79ZM64 208H40v-16h24Zm128 0v-16h24v16Zm24-32H40v-48h176ZM56 152a8 8 0 0 1 8-8h16a8 8 0 0 1 0 16H64a8 8 0 0 1-8-8Zm112 0a8 8 0 0 1 8-8h16a8 8 0 0 1 0 16h-16a8 8 0 0 1-8-8Z"/></g>`),
		g.Group(children),
	)
}