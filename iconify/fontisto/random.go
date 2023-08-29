package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Random(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24.983 8.539V6.054h-4.902l-3.672 5.945l-2.099 3.414l-3.24 5.256c-.326.51-.889.844-1.53.845H0v-3.568h8.538L12.211 12l2.099-3.414l3.24-5.256a1.812 1.812 0 0 1 1.525-.845h5.904V0l7.417 4.27l-7.417 4.27z"/><path fill="currentColor" d="m12.902 6.316l-.63 1.022l-1.468 2.39l-2.265-3.675H.001V2.485h9.54a1.813 1.813 0 0 1 1.526.838l.004.007l1.836 2.985zM24.983 24v-2.485h-5.904a1.809 1.809 0 0 1-1.521-.838l-.004-.007l-1.836-2.985l.63-1.022l1.468-2.39l2.264 3.675h4.902v-2.485l7.417 4.27l-7.417 4.27z"/>`),
		g.Group(children),
	)
}