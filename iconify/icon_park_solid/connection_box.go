package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectionBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSConnectionBox0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M43 4H5c-.552 0-1 .488-1 1.09v9.82c0 .602.448 1.09 1 1.09h38c.552 0 1-.488 1-1.09V5.09C44 4.489 43.552 4 43 4Zm0 28H5c-.552 0-1 .488-1 1.09v9.82c0 .602.448 1.09 1 1.09h38c.552 0 1-.488 1-1.09v-9.82c0-.602-.448-1.09-1-1.09Z"/><path stroke="#fff" stroke-linecap="round" d="M14 16v8.008l20 .01V32"/><path stroke="#000" stroke-linecap="round" d="M18 38h12M18 10h12"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSConnectionBox0)"/>`),
		g.Group(children),
	)
}