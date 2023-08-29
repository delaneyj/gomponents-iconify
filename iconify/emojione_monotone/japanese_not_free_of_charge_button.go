package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseNotFreeOfChargeButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M26.878 30.009H41.05v2.466H26.878z"/><path fill="currentColor" d="M52 2H12C6.477 2 2 6.478 2 12v40c0 5.523 4.477 10 10 10h40c5.523 0 10-4.477 10-10V12c0-5.522-4.477-10-10-10m-2 20.963H28.997a31.501 31.501 0 0 1-1.57 3.042h18.177v19.021c0 2.271-.432 3.504-2.041 4.197c-1.531.732-3.77.77-6.949.77c-.197-1.27-.824-3.195-1.414-4.428c2.002.117 4.281.117 4.947.117c.668-.039.904-.232.904-.771v-2.309H26.878v7.354h-4.634v-16.75c-1.57 1.695-3.259 3.234-5.183 4.504c-.667-1-2.158-2.656-3.062-3.504c4.162-2.733 7.42-6.775 9.774-11.242h-8.597v-4.312h10.56a37.52 37.52 0 0 0 1.533-4.659l4.827 1.078c-.393 1.193-.785 2.387-1.254 3.581H50v4.311"/><path fill="currentColor" d="M26.878 36.285H41.05v2.504H26.878z"/>`),
		g.Group(children),
	)
}