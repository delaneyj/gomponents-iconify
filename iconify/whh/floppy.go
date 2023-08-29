package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Floppy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.06 1024h-64V576q0-26-19-45t-45-19h-640q-27 0-45.5 19t-18.5 45v448h-64q-26 0-45-19t-19-45V64q0-26 19-45t45-19h832l128 128v832q0 26-19 45t-45 19zm-192-960h-640v256q0 27 18.5 45.5t45.5 18.5h512q26 0 45-18.5t19-45.5V64zm-256 64h128v192h-128V128zm-288 448h576q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-576q-13 0-22.5-9.5t-9.5-22.5t9.5-22.5t22.5-9.5zm0 192h576q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-576q-13 0-22.5-9.5t-9.5-22.5t9.5-22.5t22.5-9.5zm0 192h576q13 0 22.5 9.5t9.5 22.5t-9.5 22.5t-22.5 9.5h-576q-13 0-22.5-9.5t-9.5-22.5t9.5-22.5t22.5-9.5z"/>`),
		g.Group(children),
	)
}