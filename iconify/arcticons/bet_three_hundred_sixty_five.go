package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BetThreeHundredSixtyFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.901 36.79c.75.624 1.375.874 3 .874h.375c1.75 0 3.25-1.5 3.25-3.25h0c0-1.75-1.5-3.25-3.25-3.25h-3.375v-3.5h6.625"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="24.529" cy="34.331" r="3.333"/><path d="M27.492 28.899c-.494-.74-1.358-1.235-2.716-1.235h-.247a3.32 3.32 0 0 0-3.333 3.333v3.334"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.811 37.652a2.504 2.504 0 0 0 2.497-2.497h0a2.504 2.504 0 0 0-2.497-2.497h0a2.504 2.504 0 0 0 2.497-2.497h0a2.504 2.504 0 0 0-2.497-2.497m-4.12 9.239c.75.624 1.346.761 3.093.761h1m-4.093-9.126c.75-.624 1.499-.874 3.122-.874h.998M31.652 12.25v10.502c0 1.05.7 1.75 1.75 1.75h.525m-4.026-9.277h3.676m-6.073 7.527c-.525 1.05-1.75 1.75-2.975 1.75h0a3.51 3.51 0 0 1-3.5-3.5v-2.275a3.51 3.51 0 0 1 3.5-3.5h0a3.51 3.51 0 0 1 3.5 3.5v1.225h-7m-9.555-1.392a3.51 3.51 0 0 1 3.5-3.5h0a3.51 3.51 0 0 1 3.5 3.5v2.276a3.51 3.51 0 0 1-3.5 3.5h0a3.51 3.51 0 0 1-3.5-3.5m0 3.5v-14"/>`),
		g.Group(children),
	)
}