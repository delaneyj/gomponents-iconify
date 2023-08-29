package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CAndA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.333 26.331H26.8c-.304 0-.685-.152-.837-.456l-2.131-2.893c-.304-.38-.761-.837-.761-1.522s.533-1.218 1.294-1.218c.685 0 1.218.533 1.218 1.218s-.533 1.294-1.599 1.522c-1.141.229-1.902.838-1.902 1.903c0 .837.532 1.446 1.598 1.446c1.37 0 2.207-1.37 3.425-2.968m9.557 3.292h-7.351m-1.838 5.514l5.513-16.337L38.5 32.169m-17.973-5.61h0a5.49 5.49 0 0 1-5.513 5.514h0c-3.063 0-5.514-2.451-5.514-5.31V21.25c0-3.063 2.45-5.513 5.514-5.31h0c3.063 0 5.31 2.451 5.31 5.31h0"/>`),
		g.Group(children),
	)
}