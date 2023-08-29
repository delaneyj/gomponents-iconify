package flagpack

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 24"),
		g.Raw(`<defs><path id="flagpackTf0" fill="#fff" d="M0 0h32v24H0z"/></defs><g fill="none"><g clip-path="url(#flagpackTf2)"><use href="#flagpackTf0"/><path fill="#2E42A5" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/><mask id="flagpackTf1" width="32" height="24" x="0" y="0" maskUnits="userSpaceOnUse" style="mask-type:luminance"><path fill="#fff" fill-rule="evenodd" d="M0 0v24h32V0H0Z" clip-rule="evenodd"/></mask><g fill-rule="evenodd" clip-rule="evenodd" mask="url(#flagpackTf1)"><path fill="#F7FCFF" d="M18.789 6.2h8.048l-.784 1.504h-2.442v.808h1.872l-.762 1.4h-1.11v2.797L25.44 9.83l2.364 4.095h-.966l-.337-.449h-2.14l-1.506 2.781v.154l-.041-.077l-.042.077v-.154l-1.504-2.78h-2.14l-.338.448h-.966l2.364-4.095l1.828 2.88V7.703h-2.442L18.789 6.2Zm6.697 5.31l-.526.807h.995l-.47-.808Zm-5.346 0l.526.807h-.995l.47-.808Zm-3.826-.897l.83-.612l.832.612l-.293-1.026l.806-.648l-1.012-.022l-.332-1.014l-.333 1.014l-1.012.022l.807.648l-.293 1.026Zm12.141-.612l-.83.612l.293-1.026l-.807-.648l1.012-.022l.332-1.014l.333 1.014l1.012.022l-.807.648l.293 1.026l-.83-.612Zm-3.375 7.802l.83-.612l.831.612l-.293-1.026l.807-.648l-1.012-.022l-.332-1.013l-.333 1.013l-1.012.022l.807.648l-.294 1.026Zm-5.39-.612l-.831.612l.293-1.026l-.807-.648l1.012-.022l.333-1.013l.332 1.013l1.012.022l-.806.648l.293 1.026l-.831-.612Zm2.279 3.009l.831-.613l.831.613l-.293-1.027l.807-.648l-1.012-.022l-.333-1.013l-.332 1.013l-1.013.022l.807.648l-.293 1.027Z"/><path fill="#F50100" d="M8 0h4v10H8V0Z"/><path fill="#2E42A5" d="M0 0h4v10H0V0Z"/><path fill="#F7FCFF" d="M4 0h4v10H4V0Z"/></g></g><defs><clipPath id="flagpackTf2"><use href="#flagpackTf0"/></clipPath></defs></g>`),
		g.Group(children),
	)
}