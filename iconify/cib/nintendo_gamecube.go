package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NintendoGamecube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m9.088 20.172l6.272 3.62v-7.245l-6.277-3.62v7.245zm-2.703 1.557l8.975 5.177v5.093L1.973 24.275V8.822l4.412 2.547zM16 8.193l-6.271 3.624L16 15.437l6.271-3.62L16 8.197zm0-3.11l7.547 4.371l4.412-2.547L16 0L2.631 7.719l4.411 2.541zm9.609 16.646v-5.265l-2.697 1.557v2.151l-6.272 3.625v-7.245l13.387-7.724v15.448L16.64 32.005v-5.088z"/>`),
		g.Group(children),
	)
}