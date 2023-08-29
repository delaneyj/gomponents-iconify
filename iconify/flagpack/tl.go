package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackTl0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackTl1)"><use href="#flagpackTl0"/><path fill="#E31D1C" stroke="#F7FCFF" stroke-width="2" d="M0-1h-1v26h34V-1H0Z"/><path fill="#FECA00" fill-rule="evenodd" d="m0 0l24 12L0 24V0Z" clip-rule="evenodd"/><path fill="#272727" fill-rule="evenodd" d="m0 0l16 12L0 24V0Z" clip-rule="evenodd"/><path fill="#F7FCFF" fill-rule="evenodd" d="m6.648 14.409l-2.02 2.1l-.21-2.986l-2.576-1.586l2.686-.843l.44-2.958l1.859 2.342l2.712-.727l-1.402 2.776L9.5 15.203l-2.852-.794Z" clip-rule="evenodd"/></g><defs><clipPath id="flagpackTl1"><use href="#flagpackTl0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}