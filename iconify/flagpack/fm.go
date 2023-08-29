package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackFm0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackFm2)"><use href="#flagpackFm0"/><path fill="#63B3E1" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackFm1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill="#F7FCFF" fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackFm1)"><path d="m16 7.3l-1.763.927l.336-1.963l-1.426-1.391l1.971-.287L16 2.8l.882 1.786l1.971.287l-1.426 1.39l.336 1.964L16 7.3Zm0 12l-1.763.927l.336-1.963l-1.426-1.391l1.971-.287L16 14.8l.882 1.787l1.971.286l-1.426 1.39l.336 1.964L16 19.3Zm-6.2-5.8l-1.763.927l.336-1.963l-1.426-1.391l1.971-.287L9.8 9l.882 1.787l1.971.286l-1.426 1.39l.336 1.964L9.8 13.5Zm12.2 0l-1.763.927l.336-1.963l-1.426-1.391l1.971-.287L22 9l.882 1.787l1.971.286l-1.426 1.39l.336 1.964L22 13.5Z"/></g></g><defs><clipPath id="flagpackFm2"><use href="#flagpackFm0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}