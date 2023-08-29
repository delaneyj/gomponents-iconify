package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path d="M14 10.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Zm-13 0a1 1 0 0 1 1-1h2a1 1 0 0 1 0 2H2a1 1 0 0 1-1-1Zm15 0a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1Zm-1.5 4a1 1 0 0 1 1.414 0l1.414 1.414a1 1 0 1 1-1.414 1.414L14.5 15.914a1 1 0 0 1 0-1.414Zm-11-11a1 1 0 0 1 1.414 0l1.414 1.414a1 1 0 1 1-1.414 1.414L3.5 4.914a1 1 0 0 1 0-1.414Zm7 12.5a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0v-2a1 1 0 0 1 1-1Zm0-15a1 1 0 0 1 1 1v2a1 1 0 1 1-2 0V2a1 1 0 0 1 1-1Zm6.828 2.5a1 1 0 0 1 0 1.414l-1.414 1.414A1 1 0 1 1 14.5 4.914L15.914 3.5a1 1 0 0 1 1.414 0Zm-11 11a1 1 0 0 1 0 1.414l-1.414 1.414A1 1 0 1 1 3.5 15.914L4.914 14.5a1 1 0 0 1 1.414 0Z" opacity=".2"/><path fill-rule="evenodd" d="M10 13.5a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm0-6a2.5 2.5 0 1 1 0 5a2.5 2.5 0 0 1 0-5Z" clip-rule="evenodd"/><rect width="1" height="3" x="1" y="10.5" rx=".5" transform="rotate(-90 1 10.5)"/><rect width="1" height="3" x="16" y="10.5" rx=".5" transform="rotate(-90 16 10.5)"/><rect width="1" height="3" x="14" y="14.707" rx=".5" transform="rotate(-45 14 14.707)"/><rect width="1" height="3" x="3" y="3.707" rx=".5" transform="rotate(-45 3 3.707)"/><rect width="1" height="3" x="9.5" y="16" rx=".5"/><rect width="1" height="3" x="9.5" y="1" rx=".5"/><rect width="1" height="3" x="16.121" y="3" rx=".5" transform="rotate(45 16.121 3)"/><rect width="1" height="3" x="5.121" y="14" rx=".5" transform="rotate(45 5.121 14)"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}