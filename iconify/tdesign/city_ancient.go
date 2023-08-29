package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityAncient(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.146 1.483l.856-1.417l.857 1.419l.005.01l.024.039l.097.154a36.062 36.062 0 0 0 1.59 2.333c.49.66 1.024 1.32 1.54 1.852c.544.559.97.87 1.233.967c.337.125.554.16.652.16h1v2h-1v2h4v2h-1v9H3v-9H2v-2h4.002V9H5V7h1c.098 0 .316-.035.654-.16c.264-.099.69-.41 1.233-.968A18.871 18.871 0 0 0 9.43 4.02a36.006 36.006 0 0 0 1.686-2.487l.024-.04l.005-.01l.002-.001ZM8.002 9v2H16V9H8.002Zm6.427-2a22.373 22.373 0 0 1-1.46-1.785c-.361-.486-.691-.96-.968-1.37c-.276.41-.606.883-.967 1.369A22.383 22.383 0 0 1 9.573 7h4.856ZM5 13v7h3v-5h8v5h3v-7H5Zm9 7v-3h-4v3h4Z"/>`),
		g.Group(children),
	)
}