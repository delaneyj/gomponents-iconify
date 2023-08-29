package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyspaceAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M321 157q-31 0-54-22.5T244 80t22.5-54.5T321 3t54.5 22.5T398 80t-22.5 54.5T321 157zm-63 74q0-32-22.5-54.5T181 154t-54 22.5t-22 54.5v138h153V231zm-77-88q24 0 40.5-16.5T238 87t-16.5-39.5T181 31q-23 0-39.5 16.5T125 87t16.5 39.5T181 143zm245 132q0-43-31-74t-74-31t-74 31t-31 74q0 2 .5 4t.5 3h-1v182h210V275zM70 145q-26 0-44.5 18.5T7 208v113h126V212h-1v-4q0-26-18-44.5T70 145zm0-8q19 0 32.5-13.5T116 91t-13.5-32.5T70 45q-20 0-33.5 13.5T23 91t14 32.5T70 137z"/>`),
		g.Group(children),
	)
}