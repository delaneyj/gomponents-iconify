package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FacebookOfficial(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1451 0q35 0 60 25t25 60v1366q0 35-25 60t-60 25h-391V941h199l30-232h-229V561q0-56 23.5-84t91.5-28l122-1V241q-63-9-178-9q-136 0-217.5 80T820 538v171H620v232h200v595H85q-35 0-60-25t-25-60V85q0-35 25-60T85 0h1366z"/>`),
		g.Group(children),
	)
}