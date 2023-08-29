package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClapperboardOpenLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M4 11h12c1.886 0 2.828 0 3.414.586C20 12.172 20 13.114 20 15v1c0 2.828 0 4.243-.879 5.121C18.243 22 16.828 22 14 22h-4c-2.828 0-4.243 0-5.121-.879C4 20.243 4 18.828 4 16v-5Zm.001 0c-.484-1.806-.725-2.709-.517-3.485A3 3 0 0 1 4.26 6.17c.569-.568 1.471-.81 3.277-1.294l7.003-1.877c.673-.18 1.01-.27 1.305-.29a3 3 0 0 1 2.886 1.667c.131.265.221.602.402 1.275c.06.225.09.337.096.435a1 1 0 0 1-.555.962c-.089.044-.2.074-.425.134L4 11Z"/><path stroke-linecap="round" d="m14.7 2.941l-.637 5.348M8.422 4.624l-.638 5.347"/></g>`),
		g.Group(children),
	)
}