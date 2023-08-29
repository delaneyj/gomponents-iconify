package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RewindBackCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.25 12C1.25 6.063 6.063 1.25 12 1.25S22.75 6.063 22.75 12S17.937 22.75 12 22.75S1.25 17.937 1.25 12ZM12 2.75a9.25 9.25 0 1 0 0 18.5a9.25 9.25 0 0 0 0-18.5Zm2.483 4.77c1.158-.828 2.767 0 2.767 1.423v6.114c0 1.423-1.609 2.251-2.767 1.424l-4.28-3.057a1.75 1.75 0 0 1 0-2.848l4.28-3.057Zm1.267 1.423a.25.25 0 0 0-.395-.203l-4.28 3.057a.25.25 0 0 0 0 .406l4.28 3.057a.25.25 0 0 0 .395-.203V8.943Zm-4.14-1.307a.75.75 0 0 1-.174 1.046l-3.222 2.3a1.25 1.25 0 0 0 0 2.035l3.222 2.301a.75.75 0 1 1-.872 1.22l-3.221-2.3a2.75 2.75 0 0 1 0-4.476l3.221-2.3a.75.75 0 0 1 1.046.174Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}