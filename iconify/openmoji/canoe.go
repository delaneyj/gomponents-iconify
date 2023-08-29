package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Canoe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#61b2e4" stroke="#61b2e4" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M66.5 24.852s-1.906 14.296-11.438 14.296H16.938C7.405 39.148 5.5 24.852 5.5 24.852S9.313 27.71 36 27.71s30.5-2.86 30.5-2.86Z"/><path fill="#a57939" d="m40.484 18.304l4.331 2.5l-2.732.733l-12.5 21.65l.732 2.732l-5.5 9.527l-5.197-3l5.5-9.527l2.732-.732l12.5-21.65l.134-2.233"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m40.35 20.537l.134-2.232m.599 4.964l-11.5 19.918l.732 2.732l-5.5 9.527l-5.197-3l5.5-9.527l1.366-.366m14-24.248l3.465 2M35.45 39.148h19.612c9.532 0 11.438-14.296 11.438-14.296s-3.229 2.409-24.414 2.803m-8.228.049C9.113 27.555 5.5 24.852 5.5 24.852s1.906 14.296 11.438 14.296H27.25"/>`),
		g.Group(children),
	)
}