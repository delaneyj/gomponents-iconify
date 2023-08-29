package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M18 4a1 1 0 0 1 .783.378l.074.108l3 5a1 1 0 0 1-.032 1.078l-.08.103l-8.53 9.533a1.7 1.7 0 0 1-1.215.51c-.4 0-.785-.14-1.11-.417l-.135-.126l-8.5-9.5A1 1 0 0 1 2.083 9.6l.06-.115l3.013-5.022l.064-.09a.982.982 0 0 1 .155-.154l.089-.064l.088-.05l.05-.023l.06-.025l.109-.032l.112-.02L6 4h12zM9.114 7.943a1 1 0 0 0-1.371.343l-.6 1l-.06.116a1 1 0 0 0 .177 1.07l2 2.2l.09.088a1 1 0 0 0 1.323-.02l.087-.09a1 1 0 0 0-.02-1.323l-1.501-1.65l.218-.363l.055-.103a1 1 0 0 0-.398-1.268z"/></g>`),
		g.Group(children),
	)
}