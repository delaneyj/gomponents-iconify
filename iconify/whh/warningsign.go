package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Warningsign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 832q0 26-18.5 45T960 896H64q-27 0-45.5-19T0 832q0-17 9-32L448 64q8-18 11.5-24t12-18.5t18-17T512 0q17 0 29.5 10t19 21.5T576 64l439 736q9 15 9 32zM576 320q0-27-18.5-45.5T512 256t-45.5 18.5T448 320v192q0 26 19 45t45 19t45-19t19-45V320zm-64.5 320q-26.5 0-45 18.5t-18.5 45t19 45.5t45 19t45-19t19-45.5t-19-45t-45.5-18.5z"/>`),
		g.Group(children),
	)
}