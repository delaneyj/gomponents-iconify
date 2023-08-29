package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m212 640l623 665l-300-665H212zm812 772l349-772H675zM538 512l204-384H480L192 512h346zm675 793l623-665h-323zM683 512h682l-204-384H887zm827 0h346l-288-384h-262zm141-486l384 512q14 18 13 41.5t-17 40.5l-960 1024q-18 20-47 20t-47-20L17 620Q1 603 0 579.5T13 538L397 26q18-26 51-26h1152q33 0 51 26z"/>`),
		g.Group(children),
	)
}