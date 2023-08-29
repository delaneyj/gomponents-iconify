package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m32.773 20.22l157.23 216.14c34.636 24.228 67.18 52.126 97.608 84.167c-55.124-32.744-111.903-62.108-173.862-80.89l-26.373-7.252c-19.323-4.77-39.164-8.476-59.64-10.89l147.478 106.376c48.008 6.057 95.976 17.203 144.207 34.02c-92.163-7.09-183.867-6.92-273.713 22.448c153.47 8.53 305.213 31.958 450.104 114.453L277.505 78.07c5.71 102.37 42.343 193.67 86.375 282.31c-43.76-50.662-78.767-104.638-106.398-162.067L32.772 20.22z"/>`),
		g.Group(children),
	)
}