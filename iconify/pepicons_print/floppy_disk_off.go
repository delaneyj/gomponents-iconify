package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyDiskOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path d="M4.75 18.5h13a1 1 0 0 0 1-1V6.914a1 1 0 0 0-.293-.707l-2.414-2.414a1 1 0 0 0-.707-.293H4.75a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1Z"/><path fill-rule="evenodd" d="M16.75 16.5V7.328L14.922 5.5H5.75v11h11Zm1 2h-13a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h10.586a1 1 0 0 1 .707.293l2.414 2.414a1 1 0 0 1 .293.707V17.5a1 1 0 0 1-1 1Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M3.5 2.75a.5.5 0 0 0-.5.5v13a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5V5.664a.5.5 0 0 0-.146-.353l-2.415-2.415a.5.5 0 0 0-.353-.146H3.5Zm-1.5.5a1.5 1.5 0 0 1 1.5-1.5h10.586a1.5 1.5 0 0 1 1.06.44l2.415 2.414A1.5 1.5 0 0 1 18 5.664V16.25a1.5 1.5 0 0 1-1.5 1.5h-13a1.5 1.5 0 0 1-1.5-1.5v-13Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5.5 10.75A1.5 1.5 0 0 1 7 9.25h6a1.5 1.5 0 0 1 1.5 1.5v7h-1v-7a.5.5 0 0 0-.5-.5H7a.5.5 0 0 0-.5.5v7h-1v-7Zm.5-6a1.5 1.5 0 0 0 1.5 1.5h4a1.5 1.5 0 0 0 1.5-1.5v-2a.5.5 0 0 0-1 0v2a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5v-2a.5.5 0 0 0-1 0v2Z" clip-rule="evenodd"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}