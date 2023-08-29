package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M183.427 136.897c11.578 11.362 3.954 31.242-11.729 31.004c-15.682-.237-71.904-.237-87.357-.237c-15.453 0-23.261-18.467-11.853-29.982c2.767-2.793 6.33-6.412 10.236-10.382c12.203-12.397 27.762-28.205 32.902-32.878c6.786-6.169 18.112-7.464 25.444 0c7.331 7.464 30.779 31.112 42.357 42.475zm-14.608 14.261c2.214.011 2.735-2.24 1.173-3.795v.001l-39.486-39.31c-1.554-1.548-4.107-1.561-5.684-.014l-39.597 39.868c-1.578 1.548-1.064 2.812 1.146 2.823l82.448.427z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}