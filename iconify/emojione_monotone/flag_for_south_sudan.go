package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSouthSudan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M11.809 32.568v5.149l3.257-4.119l5.109 1.557l-3.14-4.131l3.14-4.135l-5.109 1.559l-3.257-4.121v5.109l-5.024 1.588z"/><path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m-2.828 30L11.191 49.98C6.862 44.985 4.5 38.674 4.5 32c0-6.676 2.362-12.988 6.69-17.981L29.172 32m-7.988 10.816l1.967-1.967h35.408a28.224 28.224 0 0 1-.734 1.967H21.184m1.966-19.667l-1.967-1.967h36.641c.271.643.512 1.301.734 1.967H23.15"/>`),
		g.Group(children),
	)
}