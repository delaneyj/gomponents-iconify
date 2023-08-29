package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePlayLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M221.89 117.69L54.05 21.62a12 12 0 0 0-12.13 0A11.69 11.69 0 0 0 36 31.87v192.26a11.69 11.69 0 0 0 5.92 10.21a12 12 0 0 0 12.13 0l167.77-96a11.76 11.76 0 0 0 .07-20.66Zm-52.44-20.8L144 122.34L50.4 28.75ZM44 222.33V33.67L138.34 128Zm6.4 4.92l93.6-93.59l25.45 25.45Zm167.51-95.88L176.65 155l-27-27l27-27L218 124.66a3.77 3.77 0 0 1-.07 6.71Z"/>`),
		g.Group(children),
	)
}