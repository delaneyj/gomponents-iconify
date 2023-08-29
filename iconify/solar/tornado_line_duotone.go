package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TornadoLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M12 3c-4.97 0-9 .672-9 1.5S7.03 6 12 6s9-.672 9-1.5c0-.52-1.588-.978-4-1.247"/><path d="M20 8s-2.791 1-8.276 1C6.24 9 4 8 4 8m12 6s-1.258.5-4 .5s-4-.5-4-.5" opacity=".5"/><path d="M16 16.5c-.748.255-1.648.5-3 .5s-2.342-.121-3-.245"/><path d="M11.719 19.46c.384.025.81.04 1.281.04a9.32 9.32 0 0 0 1.5-.114" opacity=".5"/><path d="M18 11s-.515 1-6 1m-6-1s.219.425 1.892.721M13 22h.5"/></g>`),
		g.Group(children),
	)
}