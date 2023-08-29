package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenMailboxWithLoweredFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M45.59 2c-6.73 0-27.188 6.636-32.596 8.431C6.137 12.381 2 18.81 2 27.087v15.175l26.834 5.273v13.324L37.224 62l6.045-5.702V43.882L62 32.789V17.965C62 7.889 55.894 2 45.59 2zM8.233 41.574l24.903-10.949v15.843L8.233 41.574zm27.085 18.275l-4.579-.623V47.908l4.302.846l.277-.164v11.259zm25.428-33.12l-6.413 3.558V21l-18.349 9.046v-3.101l24.762-11.658v11.442z"/>`),
		g.Group(children),
	)
}