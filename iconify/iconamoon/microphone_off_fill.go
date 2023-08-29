package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.953 3.409A4 4 0 0 1 16 6v3.343a1 1 0 1 1-2 0V6a2 2 0 0 0-3.524-1.295a1 1 0 0 1-1.524-1.296Zm5.336 9.465a1.01 1.01 0 0 0-.04-.04L9.72 8.306a1.006 1.006 0 0 0-.026-.026L4.707 3.293a1 1 0 0 0-1.414 1.414L8 9.414V11a4 4 0 0 0 5.351 3.766l1.51 1.51A6 6 0 0 1 6 11a1 1 0 1 0-2-.001a8.001 8.001 0 0 0 7 7.938V21a1 1 0 1 0 2 0v-2.062a7.958 7.958 0 0 0 3.32-1.204l2.973 2.973a1 1 0 0 0 1.414-1.414l-3.559-3.56a1 1 0 0 0-.034-.033l-2.825-2.826ZM19 10a1 1 0 0 1 1 1c0 .81-.12 1.593-.346 2.332a1 1 0 1 1-1.913-.582c.168-.553.259-1.14.259-1.75a1 1 0 0 1 1-1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}