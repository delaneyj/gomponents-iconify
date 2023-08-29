package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemperatureEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M112 112c0-26.5 21.5-48 48-48s48 21.5 48 48v164.5c0 17.3 7.1 31.9 15.3 42.5c10.5 13.6 16.7 30.5 16.7 49c0 44.2-35.8 80-80 80s-80-35.8-80-80c0-18.5 6.2-35.4 16.7-48.9c8.2-10.6 15.3-25.2 15.3-42.5V112zM160 0C98.1 0 48 50.2 48 112v164.5c0 .1-.1.3-.2.6c-.2.6-.8 1.6-1.7 2.8C27.2 304.2 16 334.8 16 368c0 79.5 64.5 144 144 144s144-64.5 144-144c0-33.2-11.2-63.8-30.1-88.1c-.9-1.2-1.5-2.2-1.7-2.8c-.1-.3-.2-.5-.2-.6V112C272 50.2 221.9 0 160 0zm0 416a48 48 0 1 0 0-96a48 48 0 1 0 0 96z"/>`),
		g.Group(children),
	)
}