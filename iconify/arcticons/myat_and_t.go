package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyatAndT(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.701 25.548h6.233m-3.058 9.084v-9.084m-4.891 6.018h-3.999m5.057 3.066l-3.057-9.084l-3.175 9.084m1.058-3.066h4.116M9.5 15.639a2.32 2.32 0 0 1 2.352-2.271h0a2.32 2.32 0 0 1 2.352 2.27v3.634M9.5 13.368v5.904m4.704-3.633a2.32 2.32 0 0 1 2.352-2.271h0a2.32 2.32 0 0 1 2.352 2.27v3.634m4.967.277l-2.47-6.018m4.704 0l-2.94 8.176a1.373 1.373 0 0 1-1.294.908h-.353m10.306 12.017h-.706a1.49 1.49 0 0 1-1.294-.681l-3.293-4.315a3.587 3.587 0 0 1-1.176-2.27a1.86 1.86 0 0 1 2-1.818a1.826 1.826 0 0 1 1.881 1.817c0 1.022-.823 1.93-2.47 2.271c-1.764.34-2.94 1.25-2.94 2.839c0 1.249.823 2.157 2.47 2.157c2.117 0 3.41-2.044 5.292-4.428m.675-4.656H38.5m-3.058 9.084v-9.084"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}