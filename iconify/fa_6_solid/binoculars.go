package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Binoculars(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M128 32h32c17.7 0 32 14.3 32 32v32H96V64c0-17.7 14.3-32 32-32zm64 96v320c0 17.7-14.3 32-32 32H32c-17.7 0-32-14.3-32-32v-59.1c0-34.6 9.4-68.6 27.2-98.3C40.9 267.8 49.7 242.4 53 216l7.5-60c2-16 15.6-28 31.8-28H192zm227.8 0c16.1 0 29.8 12 31.8 28l7.4 60c3.3 26.4 12.1 51.8 25.8 74.6c17.8 29.7 27.2 63.7 27.2 98.3V448c0 17.7-14.3 32-32 32H352c-17.7 0-32-14.3-32-32V128h99.8zM320 64c0-17.7 14.3-32 32-32h32c17.7 0 32 14.3 32 32v32h-96V64zm-32 64v160h-64V128h64z"/>`),
		g.Group(children),
	)
}