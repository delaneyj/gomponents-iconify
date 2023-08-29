package pepicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MotorcyclePrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.322 7.384A4.501 4.501 0 0 1 16 11.5V14a2 2 0 0 1-2 2h-.585c.055.156.085.325.085.5v2A1.5 1.5 0 0 1 12 20h-1a1.5 1.5 0 0 1-1.5-1.5v-2c0-.175.03-.344.085-.5H9a2 2 0 0 1-2-2v-2.5a4.5 4.5 0 0 1 2.678-4.116a3 3 0 1 1 3.643 0Z" clip-rule="evenodd" opacity=".8"/><path d="M12.75 12a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-3.5 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z"/><path fill-rule="evenodd" d="M10 7a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-5a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm.5 12h-1A1.5 1.5 0 0 0 8 15.5v2A1.5 1.5 0 0 0 9.5 19h1a1.5 1.5 0 0 0 1.5-1.5v-2a1.5 1.5 0 0 0-1.5-1.5ZM9 15.5a.5.5 0 0 1 .5-.5h1a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5h-1a.5.5 0 0 1-.5-.5v-2Z" clip-rule="evenodd"/><path d="M15.5 3.75a.75.75 0 0 1 0-1.5h2a.75.75 0 0 1 0 1.5h-2Zm-13 0a.75.75 0 0 1 0-1.5h2a.75.75 0 0 1 0 1.5h-2Z"/><path d="m4.106 3.364l.302-.97l3.698.742l-.302.97l-3.698-.742Zm7.788-.228l.302.97l3.698-.742l-.302-.97l-3.698.742Z"/><path fill-rule="evenodd" d="M14.5 10.5a4.5 4.5 0 1 0-9 0V13a2 2 0 0 0 2 2h5a2 2 0 0 0 2-2v-2.5Zm-8 0a3.5 3.5 0 1 1 7 0V13a1 1 0 0 1-1 1h-5a1 1 0 0 1-1-1v-2.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}