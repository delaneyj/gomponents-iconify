package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Duomobile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 18.13h5.87A5.87 5.87 0 0 1 16.24 24v0h0H4.5h0v-5.87h0Zm0 5.87h11.74v0a5.87 5.87 0 0 1-5.87 5.87H4.5h0V24h0Zm25.38-5.88v11.74h-5.87V18.12zm-5.88.01v11.75h0a5.87 5.87 0 0 1-5.87-5.87v-5.88H24Zm13.63 0h0A5.87 5.87 0 0 1 43.5 24v0h0h-11.74h0v0a5.87 5.87 0 0 1 5.87-5.87Zm0 11.74h0A5.87 5.87 0 0 1 31.76 24H43.5a5.87 5.87 0 0 1-5.87 5.87Z"/>`),
		g.Group(children),
	)
}